<template>
  <div class="content">
    <div class="wrap">
      <div class="add" v-if="addOrSearch === 'add'">
        <div class="editor" id="editorWrap">
          <div class="editor-def">
            <div class="bd">
              <input class="ipt-area" placeholder="写下你要收藏的链接…" id="editorIpt" v-if="!newShowTitle" v-model="newUrlText" maxlength="500" autocomplete="off"/>
              <input class="ipt-area" placeholder="写下你该链接的标题…" id="editorIpt" v-if="newShowTitle" v-model="newTitleText" maxlength="500" autocomplete="off"/>
              <div class="load" @click="getUrlTitle" v-if="!newShowTitle" :class="{loading: getTitleLoading}">
                <icon name="Loading"></icon>
              </div>
              <div class="cancel" @click="resetNewUrl" v-if="newShowTitle">
                <icon name="baseline-close-px"></icon>
              </div>
              <div class="meta" v-if="newShowTitle">
                <div class="tag-group fl">
                  <div class="tag" v-for="tag in newUrlNewTags" :key="tag">
                    <span>{{tag}}</span>
                    <span class="close" @click.stop="newUrlRemoveTag(tag)">
                      <icon name="baseline-close-px"></icon>
                    </span>
                  </div>
                  <div class="new" v-if="newShowNewButton" @click.stop="newAddTag">
                    <span>+ 添加</span>
                  </div>
                  <div class="input-outer" v-if="newAddTagEditing">
                    <input type="text" v-autowidth="{maxWidth: '90px', minWidth: '20px', comfortZone: 0}" v-model="newUrlNewTagText" @keyup="newAddTagSearchTyping" placeholder="填写分组名称..." class="input" maxlength="15" autocomplete="off" ref="newAddTagInput"/>
                    <span class="close" @click.stop="newCancelAddTag">
                      <icon name="baseline-close-px"></icon>
                    </span>
                    <div class="autocomplete" v-if="newUrlNewTagText.length && newUrlNewTagSearchedTags.length">
                      <ul class="list">
                        <div class="top-middle-arrow"></div>
                        <li class="item" v-for="tag in newUrlNewTagSearchedTags" :key="tag.id" @click.stop="newUrlTagSearchSelect(tag.tag_name)">{{tag.tag_name}}
                        </li>
                      </ul>
                    </div>
                  </div>
                </div>
              </div>
              <div class="meta cl" v-if="newShowTitle">
                <div class="btn-group fr" @click="collectUrl">
                  <span class="btn btn-green fl" id="btnAddComment" :class="{disable: newUrlNewTags.length === 0 || newTitleText.length === 0 || collectUrlPending}">
                    <span class="free">收藏</span>
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="search" v-if="addOrSearch === 'search'">
        <div class="editor" id="editorWrap">
          <div class="editor-def">
            <div class="bd">
              <input class="ipt-area" placeholder="写下你要搜索的内容…" @keyup="collectSearchTyping" v-model="collectSearchText" id="editorIpt" maxlength="20" autocomplete="off"/>
            </div>
          </div>
        </div>
      </div>
      <div class="filter" id="filterWrap">
        <div class="inner">
          <div class="currentTag" v-if="collectSearchText.length === 0">
            <a class="tag" @click.stop="clickTagDropDown">
              <span class="span">{{currentTag}}</span>
              <icon :name="tagDropDown? 'arrow_down':'arrow-up'" class="updown"></icon>
            </a>
            <div class="select-container" :class="{show: tagDropDown}">
              <div class="header">
                <icon v-if="!newTagEditing" name="search" class="search"></icon>
                <input v-if="!newTagEditing" @click.stop="" @keyup="tagSearchTyping" v-model="tagSearchText" class="searchInput" ref="tagSearchInput" placeholder="搜索分组..." autocomplete="off"/>
                <span v-if="!newTagEditing" @click.stop="clickAddTag" class="add">+创建</span>
                <input v-if="newTagEditing" @click.stop="" @keyup="tagAddTyping" v-model="tagAddText" ref="newTagInput" class="newTagInput" placeholder="回车创建分组..." maxlength="15" autocomplete="off"/>
                <span v-if="newTagEditing" @click.stop="exitNewTagEdit" class="close">
                  <icon name="baseline-close-px"></icon>
                </span>
              </div>
              <div class="tag-content" v-if="(tagSearchText.length === 0 && tagAddText.length === 0) || searchedTags.length">
                <ul class="list" id="tag-list" @scroll.passive="handleTagListScroll">
                  <li class="item" @click="changeTag('全部')" v-if="tagSearchText.length === 0 && tagAddText.length === 0">
                    全部
                  </li>
                  <!-- <li class="item" @click="changeTag('稍后阅读')" v-if="tagSearchText.length === 0 && tagAddText.length === 0">
                    稍后阅读
                  </li> -->
                  <li class="item" v-for="tag in tags" :key="tag.id" @click="changeTag(tag.name)" v-if="tagSearchText.length === 0 && tagAddText.length === 0">
                    {{tag.name}}
                  </li>
                  <li class="item" v-for="tag in searchedTags" :key="tag.id" @click="changeTag(tag.tag_name)" v-if="tagSearchText.length || tagAddText.length">
                    {{tag.tag_name}}
                  </li>
                </ul>
              </div>
            </div>
          </div>
          <div class="searchHeader" v-if="collectSearchText.length">
            <a class="tag">
              <span class="span">搜索结果...</span>
            </a>
          </div>
        </div>
      </div>
      <div class="collect">
        <div class="inner">
          <div class="collectlist" id="collectList" v-if="collectSearchText.length === 0" v-infinite-scroll="loadMoreCollectItems" infinite-scroll-disabled="scrollBusy" infinite-scroll-distance="10" infinite-scroll-immediate-check="false">
            <div class="item item-collect cl" v-for="collectItem in collectItems" :key="collectItem.id">
              <div class="inner ai">
                <div class="cont full">
                  <h2 class="title"><a :href="collectItem.url" target="_blank">{{collectItem.title}}</a></h2>
                  <div class="meta cl full">
                    <div class="fl full">
                      <span class="desc">{{collectItem.description}}</span>
                      <span class="site">{{collectItem.site_domain}}</span>
                      <span class="fr">{{format(collectItem.created_at)}}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="collectlist" id="collectList" v-if="collectSearchText.length" v-infinite-scroll="loadMoresSearchCollectItems" infinite-scroll-disabled="scrollBusy" infinite-scroll-distance="10" infinite-scroll-immediate-check="false">
            <div class="item item-collect cl" v-for="collectItem in searchedCollects" :key="collectItem.id">
              <div class="inner ai">
                <div class="cont full">
                  <h2 class="title"><a :href="collectItem.url" target="_blank" v-html="collectItem.title"></a></h2>
                  <div class="meta cl full">
                    <div class="fl full">
                      <span class="desc" v-html="collectItem.description"></span>
                      <span class="site">{{collectItem.site_domain}}</span>
                      <span class="fr">{{format(collectItem.created_at)}}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="noMore">
            <span>没有更多了</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import { format } from 'timeago.js';

export default {
  name: 'MyPage',
  data: function () {
    return {
      tagSearchTypeTimer: null,
      collectSearchTypeTimer: null,
      newAddtagSearchTypeTimer: null
    }
  },
  computed: {
    ...mapGetters([
      'collectItems',
      'tags',
      'nextCollectId',
      'currentTag',
      'tagDropDown',
      'scrollBusy',
      'addOrSearch',
      'searchedTags',
      'searchedCollects',
      'collectUrlPending'
    ]),
    tagAddText: {
      get() {
        return this.$store.state.myModule.tagAddText;
      },
      set(value) {
        this.$store.commit('UPDATE_TAG_ADD_TEXT', value);
      }
    },
    tagSearchText: {
      get() {
        return this.$store.state.myModule.tagSearchText;
      },
      set(value) {
        this.$store.commit('UPDATE_TAG_SEARCH_TEXT', value);
      }
    },
    collectSearchText: {
      get() {
        return this.$store.state.myModule.collectSearchText;
      },
      set(value) {
        this.$store.commit('UPDATE_COLLECT_SEARCH_TEXT', value);
      }
    },
    newUrlText: {
      get() {
        return this.$store.state.myModule.newUrlText;
      },
      set(value) {
        this.$store.commit('UPDATE_NEW_URL_TEXT', value);
      }
    },
    newTitleText: {
      get() {
        return this.$store.state.myModule.newTitleText;
      },
      set(value) {
        this.$store.commit('UPDATE_NEW_TITLE_TEXT', value);
      }
    },
    newShowTitle: {
      get() {
        return this.$store.state.myModule.newShowTitle;
      },
      set(value) {
        this.$store.commit('UPDATE_NEW_SHOW_TITLE', value);
      }
    },
    getTitleLoading: {
      get() {
        return this.$store.state.myModule.getTitleLoading;
      },
      set(value) {
        this.$store.commit('UPDATE_GET_TITLE_LOADING', value);
      }
    },
    newUrlNewTagText: {
      get() {
        return this.$store.state.myModule.newUrlNewTagText;
      },
      set(value) {
        this.$store.commit('UPDATE_NEW_URL_NEW_TAG_TEXT', value);
      }
    },
    newShowNewButton: {
      get() {
        return this.$store.state.myModule.newShowNewButton;
      },
      set(value) {
        this.$store.commit('UPDATE_NEW_SHOW_NEW_BUTTON', value);
      }
    },
    newAddTagEditing: {
      get() {
        return this.$store.state.myModule.newAddTagEditing;
      },
      set(value) {
        this.$store.commit('UPDATE_NEW_ADD_TAG_EDITING', value);
      }
    },
    newTagEditing: {
      get() {
        return this.$store.state.myModule.newTagEditing;
      },
      set(value) {
        this.$store.commit('UPDATE_NEW_TAG_EDITING', value);
      }
    },
    newUrlNewTagSearchedTags: {
      get() {
        return this.$store.state.myModule.newUrlNewTagSearchedTags;
      },
      set(value) {
        this.$store.commit('UPDATE_NEW_URL_NEW_TAG_SEARCHED_TAGS', value);
      }
    },
    newUrlNewTags: {
      get() {
        return this.$store.state.myModule.newUrlNewTags;
      },
      set(value) {
        this.$store.commit('UPDATE_NEW_URL_NEW_TAGS', value);
      }
    }
  },
  created() {
    this.updateInitialState();
  },
  mounted () {
    let that = this;
    document.body.addEventListener('click', function(){
      that.$store.dispatch('updateTagDropDown', {dropdown: false});
      that.tagAddText = '';
      that.tagSearchText = '';
    });
  },
  methods: {
    format(time) {
      return format(time, 'zh_CN');
    },
    updateInitialState() {
      let tag = this.$route.query.tag || '全部';
      if (tag === this.currentTag) {
        return;
      }
      this.$store.dispatch('changeCurrentTag', {tag: tag});
      this.$store.dispatch('getCollectItems');
      this.$store.dispatch('getTags');
    },
    loadMoreCollectItems() {
      this.$store.dispatch('getCollectItems');
    },
    loadMoreTags() {
      this.$store.dispatch('getTags');
    },
    clickTagDropDown() {
      this.$store.dispatch('updateTagDropDown', {dropdown: !this.tagDropDown});
      if (this.newTagEditing) {
        this.$nextTick(() => this.$refs.newTagInput.focus());
      } else {
        this.$nextTick(() => this.$refs.tagSearchInput.focus());
      }
    },
    changeTag(tag) {
      if (tag === this.currentTag) {
        return;
      }
      this.$router.push({ query: { tag: tag } });
      this.updateInitialState();
    },
    clickAddTag() {
      this.newTagEditing = true;
      this.$nextTick(() => this.$refs.newTagInput.focus());
    },
    exitNewTagEdit() {
      this.newTagEditing = false;
      this.tagAddText = '';
      this.tagSearchText= '';
    },
    tagSearchTyping() {
      let that = this;
      if (that.tagSearchTypeTimer) {
        clearTimeout(that.tagSearchTypeTimer);
      }
      that.tagSearchTypeTimer = setTimeout(function(){
        that.$store.dispatch('searchTag', {keyword: that.tagSearchText});
      }, 500);
    },
    tagAddTyping(event) {
      let that = this;
      if (event.key == "Enter") {
        that.$store.dispatch('createTag', {name: that.tagAddText});
        that.exitNewTagEdit();
      } else {
        if (that.tagSearchTypeTimer) {
          clearTimeout(that.tagSearchTypeTimer);
        }
        that.tagSearchTypeTimer = setTimeout(function(){
          that.$store.dispatch('searchTag', {keyword: that.tagAddText});
        }, 500);  
      }
    },
    collectSearchTyping() {
      let that = this;
      if (that.collectSearchTypeTimer) {
        clearTimeout(that.collectSearchTypeTimer);
      }
      that.collectSearchTypeTimer = setTimeout(function(){
        that.$store.dispatch('searchCollect', {keyword: that.collectSearchText});
      }, 500);
    },
    handleTagListScroll() {
      let that = this;
      const tagSelectContainer = document.querySelector('#tag-list');
      if(tagSelectContainer.scrollTop + tagSelectContainer.clientHeight >= tagSelectContainer.scrollHeight) {
        that.loadMoreTags();
      }
    },
    loadMoresSearchCollectItems() {
      this.$store.dispatch('searchCollectMore', {keyword: this.collectSearchText});
    },
    getUrlTitle() {
      this.$store.dispatch('getUrlTitle', {url: this.newUrlText});
    },
    resetNewUrl() {
      this.newUrlText = '';
      this.newTitleText = '';
      this.newShowTitle = false;
      this.getTitleLoading = false;
    },
    newAddTag() {
      this.newShowNewButton = false;
      this.newAddTagEditing = true;
      this.$nextTick(() => this.$refs.newAddTagInput.focus());
    },
    newCancelAddTag() {
      this.newUrlNewTagText = '';
      this.newUrlText = '';
      this.newShowNewButton = true;
      this.newAddTagEditing = false;
    },
    newAddTagSearchTyping(event) {
      let that = this;
      if (event.key == "Enter") {
        if (that.newUrlNewTagText.length && that.newUrlNewTags.indexOf(that.newUrlNewTagText) === -1) {
          if (that.newUrlNewTagText === '稍后阅读') {
            that.newUrlNewTags = ['稍后阅读'];
          } else {
            that.newUrlNewTags = that.newUrlNewTags.concat(that.newUrlNewTagText);  
          }
          that.newCancelAddTag();
          if (that.newUrlNewTags.length > 2 || that.newUrlNewTags.indexOf('稍后阅读') >= 0) {
            that.newShowNewButton = false;
          }
        }
      } else {
        if (that.newAddtagSearchTypeTimer) {
          clearTimeout(that.newAddtagSearchTypeTimer);
          that.newUrlNewTagSearchedTags = [];
        }
        that.newAddtagSearchTypeTimer = setTimeout(function(){
          that.$store.dispatch('newAddSearchTag', {keyword: that.newUrlNewTagText});
        }, 500);  
      }
    },
    newUrlRemoveTag(tag) {
      let index = this.newUrlNewTags.indexOf(tag);
      if (index >= 0) {
        this.newUrlNewTags.splice(index, 1);
        this.newCancelAddTag();
      }
    },
    newUrlTagSearchSelect(tag) {
      let that = this;
      if (tag.length && that.newUrlNewTags.indexOf(tag) === -1) {
        if (tag === '稍后阅读') {
          that.newUrlNewTags = ['稍后阅读'];
        } else {
          that.newUrlNewTags = that.newUrlNewTags.concat(tag);  
        }
        that.newCancelAddTag();
        if (that.newUrlNewTags.length > 2 || that.newUrlNewTags.indexOf('稍后阅读') >= 0) {
          that.newShowNewButton = false;
        }
      }
    },
    collectUrl() {
      if (this.newUrlNewTags.length && this.newTitleText.length && !this.collectUrlPending) {
        this.$store.dispatch('collectUrl');
      }
    }
  },
  watch: {
    $route(to, from) {
      if (from.path === '/my' && to.path === '/my') {
        let fromTag = from.query.tag || '全部';
        let toTag = to.query.tag || '全部';
        if (toTag === fromTag) {
          return;
        }
        this.$store.dispatch('changeCurrentTag', {tag: toTag});
        this.$store.dispatch('getCollectItems');
      }
    }
  }
};
</script>

<style scoped>
.content {
    padding-top: 60px;
}

.content .wrap {
    position: relative;
    width: 80%;
    margin: 30px auto 0;
}

.content .add .editor .editor-def .bd {
    margin-top: 10px;
    padding: 20px;
    border: 1px solid #F1F1F1;
    -webkit-box-shadow: 0 5px 20px 0 rgba(0, 0, 0, 0.05);
    box-shadow: 0 5px 20px 0 rgba(0, 0, 0, 0.05);
    border-radius: 4px;
    background: #fff;
    position: relative;
}

.content .add .editor-focus .editor-def .bd .ipt-area {
    height: 72px;
}

.content .add .editor .editor-def .bd .ipt-area {
    width: calc(100% - 30px);
    height: 24px;
    padding: 2px 0;
    line-height: 24px;
    resize: unset;
    border: none;
    margin-right: 10px;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
}

.content .add .editor .editor-def .bd .load {
    display: inline-block;
    width: 16px;
    height: 16px;
    line-height: 16px;
    font-size: 16px;
    cursor: pointer;
    vertical-align: middle;
}

.content .add .editor .editor-def .bd .cancel {
    display: inline-block;
    width: 16px;
    height: 16px;
    line-height: 16px;
    font-size: 16px;
    cursor: pointer;
    vertical-align: middle;
}

.content .add .editor .editor-def .bd .load.loading {
    -webkit-animation: rotation 2s infinite linear;
}

@-webkit-keyframes rotation {
    from {-webkit-transform: rotate(0deg);}
    to   {-webkit-transform: rotate(359deg);}
}

.content .add .editor .editor-def .bd .load.hide {
    display: none;
}

.content .add .editor .editor-def .bd .meta .tag-group {
    margin-top: 1px;
}

.content .add .editor .editor-def .bd .meta .tag-group .tag {
    height: 20px;
    margin-right: 8px;
    padding: 0 10px;
    line-height: 22px;
    background-color: #fff;
    border: 1px solid #5ed5c8;
    color: #70797b;
    font-size: 12px;
    border-radius: 10px;
    display: inline-block;
}

.content .add .editor .editor-def .bd .meta .tag-group .new {
    height: 20px;
    margin-right: 8px;
    padding: 0 10px;
    line-height: 22px;
    background-color: #fff;
    border: 1px solid #5ed5c8;
    color: #70797b;
    font-size: 12px;
    border-radius: 10px;
    display: inline-block;
    cursor: pointer;
}

.content .add .editor .editor-def .bd .meta .tag-group .close {
    cursor: pointer;
    margin-left: 3px;
}

.content .add .editor .editor-def .bd .meta .tag-group .close:hover {
    color: #fff;
    background-color: #5ed5c8;
    border-radius: 50%;
}

.content .add .editor .editor-def .bd .meta .tag-group .new:hover {
    color: #fff;
    background-color: #5ed5c8;
}

.content .add .editor .editor-def .bd .meta .tag-group .input-outer {
    height: 20px;
    margin-right: 8px;
    padding: 0 10px;
    line-height: 22px;
    background-color: #fff;
    border: 1px solid #5ed5c8;
    color: #70797b;
    font-size: 12px;
    border-radius: 10px;
    display: inline-block;
    position: relative;
}

.content .add .editor .editor-def .bd .meta .tag-group .input-outer .input {
    width: 80px;
    height: 12px;
    padding: 2px 0;
    line-height: 12px;
    resize: unset;
    border: none;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
    color: #70797b;
}

.content .add .editor .editor-def .bd .meta .tag-group .input-outer .autocomplete {
    position: absolute;
    top: 20px;
    left: 50%;
    transform: translateX(-50%);
    margin: 12px 0 10px 0;
    padding: 6px 0 6px 0;
    background-color: #fff;
    border: 1px solid #ebeef5;
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0,0,0,.1);
    width: 90px;
    z-index: 1;
}

.content .add .editor .editor-def .bd .meta .tag-group .input-outer .top-middle-arrow {
    position: absolute;
    top: -12px;
    left: 50%;
    transform: translateX(-50%);
    filter: drop-shadow(0 2px 12px rgba(0,0,0,.03));
    border-style: solid;
    border-color: transparent;
    border-top-width: 0;
    border-bottom-color: #ebeef5;
    margin-right: 3px;
    border-width: 6px;
    position: absolute;
    display: block;
    width: 0;
    height: 0;   
}

.content .add .editor .editor-def .bd .meta .tag-group .input-outer .top-middle-arrow:after {
    position: absolute;
    display: block;
    width: 0;
    height: 0;
    border-color: transparent;
    border-style: solid;
    content: " ";
    border-width: 6px;
    top: 1px;
    margin-left: -6px;
    border-top-width: 0;
    border-bottom-color: #fff;
}

.content .add .editor .editor-def .bd .meta .tag-group .input-outer .list .item {
    list-style: none;
    line-height: 24px;
    padding: 0 8px;
    margin: 0;
    font-size: 12px;
    color: #606266;
    cursor: pointer;
    outline: none;
}

.content .add .editor .editor-def .bd .meta .tag-group .input-outer .list .item:hover {
    background-color: rgba(9,30,66,.04);
    color: #172b4d;
}

.content .add .editor .editor-def .bd .meta {
    display: block;
    margin-top: 22px;
    line-height: 34px;
}

.content .add .editor .editor-def .bd .meta .btn {
    width: auto;
    height: 34px;
    padding: 0 14px;
    line-height: 34px;
    border-radius: 3px;
    display: inline-block;
    text-align: center;
    color: #fff;
    cursor: pointer;
    box-sizing: border-box;
}

.content .add .editor .editor-def .bd .meta .btn-text {
    margin-left: 20px;
    line-height: 34px;
    text-decoration: none;
    cursor: pointer;
    display: inline-block;
    color: #82B64A;
}

.content .search .editor .editor-def .bd {
    margin-top: 10px;
    padding: 20px;
    border: 1px solid #F1F1F1;
    -webkit-box-shadow: 0 5px 20px 0 rgba(0, 0, 0, 0.05);
    box-shadow: 0 5px 20px 0 rgba(0, 0, 0, 0.05);
    border-radius: 4px;
    background: #fff;
}

.content .search .editor-focus .editor-def .bd .ipt-area {
    height: 72px;
}

.content .search .editor .editor-def .bd .ipt-area {
    display: block;
    width: 100%;
    height: 24px;
    padding: 2px 0;
    line-height: 24px;
    resize: unset;
    border: none;
}

.filter {
    margin-top: 20px;
}

.filter .inner .currentTag {
    margin-right: 20px;
    font-size: 18px;
    position: relative;
}

.filter .inner .searchHeader {
    margin-right: 20px;
    font-size: 18px;
    position: relative;
    border-bottom: 1px solid #f7f7f7;
}

.filter .inner .searchHeader .tag .span {
    line-height: 48px;
    height: 48px;
}

.filter .inner .searchHeader a {
    color: black;
    font-weight: 700;
    display: inline-block;
}

.filter .inner .currentTag .tag .span {
    line-height: 48px;
    height: 48px;
    cursor: pointer;
}

.filter .inner .currentTag .tag .updown {
    line-height: 48px;
    height: 48px;
    font-size: 12px;
    cursor: pointer;
    vertical-align: top;
}

.filter .inner .currentTag .select-container.show {
    display: block;
}

.filter .inner .currentTag a {
    color: black;
    font-weight: 700;
    border-bottom: 1px solid #151515;
    display: inline-block;
}

a {
    text-decoration: none;
}

.filter .inner .currentTag .select-container {
    position: absolute;
    display: none;
    top: 45px;
    left: -5px;
    width: 270px;
    background-color: white;
    border: 1px solid #f1f1f1;
    -webkit-box-shadow: 0 5px 20px 0 rgba(0,0,0,.05);
    box-shadow: 0 5px 20px 0 rgba(0,0,0,.05);
    border-radius: 4px;
    z-index: 1;
}

.filter .inner .currentTag .select-container .header {
    padding:0 12px;
    position: relative;
    border-bottom: 1px solid rgba(9,30,66,.13);
}

.filter .inner .currentTag .select-container .search {
    color: #6b778c;
    padding: 10px 12px 10px 8px;
    position: absolute;
    margin: 0;
    left: 12px;
    top: 6px;
    z-index: 1;
    font-size: 16px;
}

.filter .inner .currentTag .select-container .close {
    color: #999;
    font-size: 16px;
    display: inline-block;
    cursor: pointer;
}

.filter .inner .currentTag .select-container .add {
    color: #5e6c84;
    font-size: 14px;
    font-family: SFProText-Regular, "Microsoft YaHei", "Open Sans", sans-serif, "Hiragino Sans GB", Arial, "Lantinghei SC", "5FAE8F6F96C59ED1", "STHeiti", "WenQuanYi Micro Hei", SimSun;
    display: inline-block;
    margin-left: 26px;
    cursor: pointer;
}

.filter .inner .currentTag .select-container .add:hover {
    color: #172b4d;
}

.filter .inner .currentTag .select-container .searchInput {
    box-sizing: border-box;
    background-color: rgb(245, 245, 245);
    outline: none;
    color: #5e6c84;
    display: block;
    height: 30px;
    line-height: 30px;
    width: 179px;
    font-family: SFProText-Regular, "Microsoft YaHei", "Open Sans", sans-serif, "Hiragino Sans GB", Arial, "Lantinghei SC", "5FAE8F6F96C59ED1", "STHeiti", "WenQuanYi Micro Hei", SimSun;
    margin-top: 8px;
    margin-bottom: 8px;
    font-size: 14px;
    overflow: hidden;
    padding: 0 10px 0 30px;
    position: relative;
    text-overflow: ellipsis;
    border-width: initial;
    border-style: none;
    border-color: initial;
    display: inline-block;
    border-radius: 20px;
    white-space: nowrap;
}

.filter .inner .currentTag .select-container .newTagInput {
    box-sizing: border-box;
    outline: none;
    color: #494949;
    display: block;
    height: 30px;
    line-height: 30px;
    width: 230px;
    font-family: SFProText-Regular, "Microsoft YaHei", "Open Sans", sans-serif, "Hiragino Sans GB", Arial, "Lantinghei SC", "5FAE8F6F96C59ED1", "STHeiti", "WenQuanYi Micro Hei", SimSun;
    margin-top: 8px;
    margin-bottom: 8px;
    font-size: 14px;
    overflow: hidden;
    padding: 0 10px 0 10px;
    position: relative;
    text-overflow: ellipsis;
    border-width: initial;
    border-style: none;
    border-color: initial;
    display: inline-block;
    white-space: nowrap;
}

.filter .inner .currentTag .select-container .tag-content {
    padding-top: 8px;
    padding-bottom: 8px;
}

.filter .inner .currentTag .select-container .tag-content .list {
    list-style: none;
    margin: 0 4px;
    padding: 0 4px;
    overflow-y: auto;
    max-height: 390px;
}

.filter .inner .currentTag .select-container .tag-content .list::-webkit-scrollbar {
    height: 8px;
    width: 8px;
}

.filter .inner .currentTag .select-container .tag-content .list::-webkit-scrollbar-thumb {
    background: rgba(9,30,66,.15);
}

.filter .inner .currentTag .select-container .tag-content .list .item {
    background-color: transparent;
    border: none;
    background: #fff;
    border-radius: 0;
    box-shadow: none;
    color: #494949;
    display: block;
    height: 100%;
    padding: 6px 0 6px 16px;
    text-align: left;
    text-decoration: none;
    transition: none;
    margin: 0;
    outline: 0;
    font-weight: 400;
    line-height: 20px;
    font-size: 14px;
    cursor: pointer;
    font-family: SFProText-Regular, "Microsoft YaHei", "Open Sans", sans-serif, "Hiragino Sans GB", Arial, "Lantinghei SC", "5FAE8F6F96C59ED1", "STHeiti", "WenQuanYi Micro Hei", SimSun;
}

.filter .inner .currentTag .select-container .tag-content .list .item:hover {
    background-color: rgba(9,30,66,.04);
    color: #172b4d;
}

.collectlist {
    margin-top: 1px;
}

.collectlist .item {
    position: relative;
    border-bottom: 1px solid #f7f7f7;
}

.collectlist .item .ai {
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    align-items: center;
    min-height: 80px;
}

.collectlist .item .inner:hover {
    background: #fafafa;
    border-radius: 4px;
}

.collectlist .item .inner {
    padding: 20px;
    margin: 0 -20px;
}

.collectlist .item .title {
    font-weight: 700;
    font-size: 18px;
    line-height: 28px;
}

.collectlist .item .title /deep/ em {
    font-weight: 700;
    font-size: 18px;
    line-height: 28px;
    color: #c00;
}

.collectlist .item .title a {
    display: block;
    color: #151515;
}

.collectlist .item .meta {
    height: 24px;
    line-height: 24px;
    color: #999;
    font-size: 0;
}

.collectlist .item .meta a, .collectlist .item .meta span {
    display: inline-block;
    font-size: 12px;
    margin-right: 10px;
    vertical-align: middle;
}

.collectlist .item .meta a, .collectlist .item .meta span.desc {
    color: #494949;
    display: block;
}

.collectlist .item .meta span.desc /deep/ em {
    color: #c00;
    font-weight: 600;
}

.collectlist .item .meta a, .collectlist .item .meta span.site {
    color: #006621;
}

.content .collect .noMore {
    margin-top: 20px;
    margin-bottom: 60px;
    text-align: center;
}

.content .collect .noMore span {
    color: #999;
    font-size: 12px;
}

.btn-green {
    background-color: #82b64a;
}

.btn-green.disable {
    background: rgb(192, 218, 164);
}
</style>
