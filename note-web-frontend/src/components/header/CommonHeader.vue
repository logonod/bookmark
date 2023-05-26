<template>
  <div id="Header" class="header cl">
    <div class="header-container cl">
      <div class="header-left fl">
        <h1 class="logo fl">
          <a href="./">
            <img src="../../assets/images/logo.png">
          </a>
        </h1>
      </div>
      <div class="header-right fr">
        <ul class="nav">
          <!-- <li class="fl"><a href="javascript:;" target="_blank">APP 下载</a></li> -->
          <li class="fl">
            <a @click="doAdd" href="javascript:;">
              <icon name="add" style="position: relative; top: -1px;"></icon>
            </a>
          </li>
          <li class="fl">
            <a @click="doSearch" href="javascript:;">
              <icon name="search"></icon>
            </a>
          </li>
          <!-- <li class="notices fl">
            <a href="javascript:;">
              <icon name="notice"></icon>
            </a>
          </li> -->
          <li class="user dropdown fl">
            <img class="avatar" src="../../assets/images/avatar.png"> 
            <div class="menu">
              <ul>
                <li class="name"><span>访客用户</span></li>
                <li><a href="https://github.com/logonod/bookmark">Github</a></li>
                <!-- <li id="inviteCode"><a href="javascript:;">邀请码</a></li>
                <li><a href="javascript:;">我的推荐</a></li>
                <li><a href="javascript:;">我的收藏</a></li>
                <li><a href="javascript:;">我的资产</a></li>
                <li><a href="javascript:;">个人设置</a></li>
                <li><a href="javascript:;">建议反馈</a></li>
                <li><a href="javascript:;">退出登录</a></li> -->
              </ul>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import Driver from "driver.js";
import "driver.js/dist/driver.min.css";

export default {
  name: 'CommonHeader',
  methods: {
    doAdd() {
      this.$store.dispatch('updateAddOrSearch', {addOrSearch: 'add'});
      setTimeout(() => {
        const driver = new Driver();
        driver.highlight({
          element: '#editorWrap',
          popover: {
            title: '提示',
            description: '输入链接再点击右侧的加载按钮，回车进入下一步，链接收藏后会进行爬取全文检索，等待后刷新页面即可看到爬取结果',
            position: 'bottom',
            offset: 20
          }
        });  
      });
    },
    doSearch() {
      this.$store.dispatch('updateAddOrSearch', {addOrSearch: 'search'});
      setTimeout(() => {
        const driver = new Driver();
        driver.highlight({
          element: '#editorWrap',
          popover: {
            title: '提示',
            description: '输入待搜索的关键字',
            position: 'bottom',
            offset: 20
          }
        });  
      });
    }
  }
};
</script>

<style scoped>
#Header {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 60px;
    padding: 0 30px;
    background: #fff;
    box-shadow: 0 0 10px 0 rgba(0,0,0,.1);
    color: #151515;
    font-size: 14px;
    line-height: 1.5;
    font-family: Helvetica Neue,Helvetica,PingFang SC,Hiragino Sans GB,Microsoft YaHei,Noto Sans CJK SC,WenQuanYi Micro Hei,Arial,sans-serif;
    box-sizing: border-box;
    z-index: 100;
}

#Header a, #Header article, #Header aside, #Header b, #Header button, #Header dd, #Header div, #Header dl, #Header dt, #Header figcaption, #Header figure, #Header footer, #Header form, #Header h1, #Header h2, #Header h3, #Header h4, #Header h5, #Header h6, #Header header, #Header i, #Header input, #Header li, #Header nav, #Header p, #Header section, #Header select, #Header span, #Header textarea, #Header ul {
    padding: 0;
    margin: 0;
    list-style: none;
    font-style: normal;
    text-decoration: none;
    border: none;
    box-sizing: border-box;
    -webkit-tap-highlight-color: transparent;
    -webkit-font-smoothing: antialiased;
}

#Header .header-left {
  width: 50%;
}

#Header .header-left, #Header .header-right {
    position: relative;
    z-index: 1;
    height: 60px;
}

#Header .fl {
    float: left;
}

#Header .logo {
    height: 100%;
}

#Header .logo a {
    display: block;
    height: 100%;
}

#Header a {
    color: #151515;
}

#Header .logo img {
    height: 60px;
    vertical-align: top;
}

#Header .header-left, #Header .header-right {
    position: relative;
    z-index: 1;
    height: 60px;
}

#Header .fr {
    float: right;
}

#Header .nav {
    display: inline-block;
}

#Header .nav li {
    padding: 0 12px;
    height: 100%;
    line-height: 60px;
}

#Header .nav li a {
    display: block;
    height: 60px;
    line-height: 60px;
}

#Header .nav li a:hover {
    color: #666;
}

#Header .nav .dropdown {
    position: relative;
}

#Header .nav .user {
    height: 60px;
    padding-right: 0;
}

#Header .nav .user .avatar {
    width: 30px;
    height: 30px;
    border-radius: 100%;
    vertical-align: middle;
}

#Header .nav .dropdown:hover .menu {
    display: block;
}

#Header .nav .dropdown .menu {
    display: none;
    position: absolute;
    top: 50px;
    right: -20px;
    width: 200px;
    padding: 10px 0;
    background: #fff;
    border: 1px solid #f1f1f1;
    box-shadow: 0 5px 20px 0 rgba(0,0,0,.05);
    border-radius: 4px;
}

#Header .nav .dropdown .menu li {
    height: 40px;
    padding: 0;
}

#Header .nav .dropdown .menu .name span {
    height: 40px;
    margin: 0;
    padding: 0 20px;
    line-height: 40px;
    font-size: 18px;
    font-weight: 700;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    display: block;
    width: 100%;
}

#Header .nav .dropdown .menu li a {
    height: 40px;
    margin: 0;
    padding: 0 20px;
    line-height: 40px;
}

#Header .nav .dropdown .menu li a:hover {
    background: #fafafa;
    color: #151515;
}
</style>
