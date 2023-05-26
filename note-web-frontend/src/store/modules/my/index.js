import axios from 'axios';
import qs from 'qs';

const state = {
  collectItems: [],
  nextCollectId: null,
  tags: [],
  nextTagId: null,
  currentTag: null,
  tagDropDown: false,
  scrollBusy: false,
  addOrSearch: 'search',
  searchedTags: [],
  searchedCollects: [],
  nextCollectSearchPage: 0,
  tagAddText: '',
  tagSearchText: '',
  collectSearchText: '',
  newUrlText: '',
  newTitleText: '',
  newShowTitle: false,
  getTitleLoading: false,
  newUrlNewTagText: '',
  newShowNewButton: true,
  newUrlNewTagSearchedTags: [],
  newUrlNewTags: [],
  newAddTagEditing: false,
  newTagEditing: false,
  collectUrlPending: false,
  newUrl: ''
}

const mutations = {
  UPDATE_COLLECT_ITEMS (state, payload) {
    state.collectItems = payload;
  },
  UPDATE_NEXT_COLLECT_ID (state, payload) {
    state.nextCollectId = payload;
  },
  PREPEND_COLLECT_ITEMS (state, payload) {
    state.collectItems = [payload, ...state.collectItems];
  },
  APPEND_COLLECT_ITEMS (state, payload) {
    state.collectItems = state.collectItems.concat(payload);
  },
  UPDATE_TAGS (state, payload) {
    state.tags = payload;
  },
  UPDATE_NEXT_TAG_ID (state, payload) {
    state.nextTagId = payload;
  },
  APPEND_TAGS (state, payload) {
    state.tags = state.tags.concat(payload);
  },
  PREAPPEND_TAG (state, payload) {
    state.tags = [payload, ...state.tags]
  },
  UPDATE_COLLECT_TAG (state, payload) {
    state.currentTag = payload;
  },
  UPDATE_TAG_DROPDOWN (state, payload) {
    state.tagDropDown = payload;
  },
  UPDATE_SCROLL_BUSY (state, payload) {
    state.scrollBusy = payload;
  },
  UPDATE_ADD_OR_SEARCH (state, payload) {
    state.addOrSearch = payload;
  },
  UPDATE_SEARCHED_TAG (state, payload) {
    state.searchedTags = payload;
  },
  UPDATE_NEXT_COLLECT_SEARCH_PAGE (state, payload) {
    state.nextCollectSearchPage = payload;
  },
  APPEND_SEARCHED_COLLECT (state, payload) {
    state.searchedCollects = state.searchedCollects.concat(payload);
  },
  UPDATE_SEARCHED_COLLECT (state, payload) {
    state.searchedCollects = payload;
  },
  UPDATE_TAG_ADD_TEXT (state, payload) {
    state.tagAddText = payload;
  },
  UPDATE_TAG_SEARCH_TEXT (state, payload) {
    state.tagSearchText = payload;
  },
  UPDATE_COLLECT_SEARCH_TEXT (state, payload) {
    state.collectSearchText = payload;
  },
  UPDATE_NEW_URL_TEXT (state, payload) {
    state.newUrlText = payload;
  },
  UPDATE_NEW_TITLE_TEXT (state, payload) {
    state.newTitleText = payload;
  },
  UPDATE_NEW_SHOW_TITLE (state, payload) {
    state.newShowTitle = payload;
  },
  UPDATE_GET_TITLE_LOADING (state, payload) {
    state.getTitleLoading = payload;
  },
  UPDATE_NEW_URL_NEW_TAG_TEXT (state, payload) {
    state.newUrlNewTagText = payload;
  },
  UPDATE_NEW_SHOW_NEW_BUTTON (state, payload) {
    state.newShowNewButton = payload;
  },
  UPDATE_NEW_URL_NEW_TAGS (state, payload) {
    state.newUrlNewTags = payload;
  },
  UPDATE_NEW_ADD_TAG_EDITING (state, payload) {
    state.newAddTagEditing = payload;
  },
  UPDATE_NEW_TAG_EDITING (state, payload) {
    state.newTagEditing = payload;
  },
  UPDATE_NEW_URL_NEW_TAG_SEARCHED_TAGS (state, payload) {
    state.newUrlNewTagSearchedTags = payload;
  },
  UPDATE_COLLECT_URL_PENDING (state, payload) {
    state.collectUrlPending = payload;
  },
  UPDATE_NEW_URL (state, payload) {
    state.newUrl = payload;
  }
}

const actions = {
  getCollectItems ({ commit }) {
    let query = {tag_name: state.currentTag};
    if (state.nextCollectId !== '') {
      commit('UPDATE_SCROLL_BUSY', true);
      query.next_collectid = state.nextCollectId;
      axios.get('/api/user/collect/list', { params: query }).then((response) => {
        commit('APPEND_COLLECT_ITEMS', response.data.collects);
        commit('UPDATE_NEXT_COLLECT_ID', response.data.next_collectid);
        commit('UPDATE_SCROLL_BUSY', false);
      });
    }
  },
  getTags ({ commit }) {
    let query = {};
    if (state.nextTagId !== '') {
      query.next_tagid = state.nextTagId;
      axios.get('/api/user/tag/list', { params: query }).then((response) => {
        commit('APPEND_TAGS', response.data.tags);
        commit('UPDATE_NEXT_TAG_ID', response.data.next_tagid);
      });
    }
  },
  changeCurrentTag ({ commit }, {tag}) {
    if (tag !== state.currentTag) {
      commit('UPDATE_COLLECT_TAG', tag);
      commit('UPDATE_NEXT_COLLECT_ID', null);
      commit('UPDATE_COLLECT_ITEMS', []);
    }
  },
  updateTagDropDown ({ commit }, {dropdown}) {
    commit('UPDATE_TAG_DROPDOWN', dropdown);
  },
  updateAddOrSearch ({ commit }, {addOrSearch}) {
    commit('UPDATE_ADD_OR_SEARCH', addOrSearch);
  },
  searchTag ({ commit }, {keyword}) {
    if (keyword) {
      let query = {keyword: keyword};
      axios.get('/api/user/tag/search', { params: query }).then((response) => {
        commit('UPDATE_SEARCHED_TAG', response.data.tags);
      });  
    }
  },
  newAddSearchTag ({ commit }, {keyword}) {
    if (keyword) {
      let query = {keyword: keyword};
      axios.get('/api/user/tag/search', { params: query }).then((response) => {
        commit('UPDATE_NEW_URL_NEW_TAG_SEARCHED_TAGS', response.data.tags);
      });  
    }
  },
  createTag ({ commit }, {name}) {
    if (name) {
      let query = {name: name};
      axios.post('/api/user/tag/create', null, { params: query }).then((response) => {
        commit('PREAPPEND_TAG', response.data.tag);
      });  
    }
  },
  searchCollect ({ commit }, {keyword}) {
    if (keyword) {
      commit('UPDATE_NEXT_COLLECT_SEARCH_PAGE', 0);
      let query = {keyword: keyword, page: state.nextCollectSearchPage};
      axios.get('/api/user/collect/search', { params: query }).then((response) => {
        commit('UPDATE_SEARCHED_COLLECT', response.data.collects);
        if (response.data.collects.length > 0) {
          commit('UPDATE_NEXT_COLLECT_SEARCH_PAGE', 1);
        } else {
          commit('UPDATE_NEXT_COLLECT_SEARCH_PAGE', -1);
        }
      });  
    }
  },
  searchCollectMore ({ commit }, {keyword}) {
    if (keyword && state.nextCollectSearchPage !== -1) {
      commit('UPDATE_SCROLL_BUSY', true);
      let query = {keyword: keyword, page: state.nextCollectSearchPage};
      axios.get('/api/user/collect/search', { params: query }).then((response) => {
        if (response.data.collects.length > 0) {
          commit('APPEND_SEARCHED_COLLECT', response.data.collects);
          commit('UPDATE_NEXT_COLLECT_SEARCH_PAGE', state.nextCollectSearchPage + 1);
        } else {
          commit('UPDATE_NEXT_COLLECT_SEARCH_PAGE', -1);
        }
        commit('UPDATE_SCROLL_BUSY', false);
      });  
    }
  },
  getUrlTitle ({ commit }, {url}) {
    let query = {url: decodeURIComponent(url)};
    if (!state.getTitleLoading && url.length) {
      commit('UPDATE_GET_TITLE_LOADING', true);
      axios.get('/api/link/get_title', { params: query }).then((response) => {
        if (response.data.errcode) {
          commit('UPDATE_GET_TITLE_LOADING', false);
        } else {
          commit('UPDATE_NEW_URL', decodeURIComponent(url));
          commit('UPDATE_NEW_URL_TEXT', '');
          commit('UPDATE_NEW_TITLE_TEXT', response.data.title);
          commit('UPDATE_NEW_SHOW_TITLE', true);
          commit('UPDATE_GET_TITLE_LOADING', false);
          commit('UPDATE_NEW_TAG_EDITING', false);
          commit('UPDATE_NEW_SHOW_NEW_BUTTON', true); 
          if (state.currentTag === '全部') {
            commit('UPDATE_NEW_URL_NEW_TAGS', []);
          } else {
            commit('UPDATE_NEW_URL_NEW_TAGS', [state.currentTag]);
          }
          if (state.currentTag === '稍后阅读') {
            commit('UPDATE_NEW_SHOW_NEW_BUTTON', false); 
          }
        }
      });  
    }
  },
  collectUrl ({ commit, dispatch }) {
    if (state.newUrl.length && state.newUrlNewTags.length && state.newTitleText.length) {
      commit('UPDATE_COLLECT_URL_PENDING', true); 
      let query = {
        u: decodeURIComponent(state.newUrl),
        t: state.newTitleText,
        n: state.newUrlNewTags
      };
      axios.post('/api/user/collect/create', null, { params: query, paramsSerializer: params => qs.stringify(params, {arrayFormat: 'repeat'}) }).then((response) => {
        commit('UPDATE_COLLECT_URL_PENDING', false);
        if (state.currentTag === '全部' || state.newUrlNewTags.indexOf(state.currentTag) >= 0) {
          commit('PREPEND_COLLECT_ITEMS', response.data.collect);
        }
        commit('UPDATE_NEW_URL_TEXT', '');
        commit('UPDATE_NEW_TITLE_TEXT', '');
        commit('UPDATE_NEW_SHOW_TITLE', false);
        commit('UPDATE_GET_TITLE_LOADING', false);
        commit('UPDATE_COLLECT_URL_PENDING', false);
        commit('UPDATE_NEXT_TAG_ID', null);
        commit('UPDATE_TAGS', []);
        dispatch('getTags');
      });
    }
  }
}

const getters = {
  collectItems: state => state.collectItems,
  nextCollectId: state => state.nextCollectId,
  tags: state => state.tags,
  nextTagId: state => state.nextTagId,
  currentTag: state => state.currentTag,
  tagDropDown: state => state.tagDropDown,
  scrollBusy: state => state.scrollBusy,
  addOrSearch: state => state.addOrSearch,
  searchedTags: state => state.searchedTags,
  nextCollectSearchPage: state => state.nextCollectSearchPage,
  searchedCollects: state => state.searchedCollects,
  tagAddText: state => state.tagAddText,
  tagSearchText: state => state.tagSearchText,
  collectSearchText: state => state.collectSearchText,
  newUrlText: state => state.newUrlText,
  newTitleText: state => state.newTitleText,
  newShowTitle: state => state.newShowTitle,
  getTitleLoading: state => state.getTitleLoading,
  newUrlNewTagText: state => state.newUrlNewTagText,
  newShowNewButton: state => state.newShowNewButton,
  newUrlNewTagSearchedTags: state => state.newUrlNewTagSearchedTags,
  newUrlNewTags: state => state.newUrlNewTags,
  newAddTagEditing: state => state.newAddTagEditing,
  newTagEditing: state => state.newTagEditing,
  collectUrlPending: state => state.collectUrlPending,
  newUrl: state => state.newUrl
}

const myModule = {
  state,
  mutations,
  actions,
  getters
}

export default myModule;
