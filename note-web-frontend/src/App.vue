<template>
  <div id="app">
    <CommonHeader />
    <router-view></router-view>
  </div>
</template>

<script>
import CommonHeader from './components/header/CommonHeader.vue';

export default {
  name: 'App',
  components: {
    CommonHeader
  }
}
</script>

<style>
html, body {
  height: 100%;
}

blockquote,
body,
button,
code,
dd,
div,
dl,
dt,
fieldset,
form,
h1,
h2,
h3,
h4,
h5,
h6,
input,
legend,
li,
ol,
optgroup,
option,
p,
pre,
select,
td,
textarea,
th,
ul {
  margin: 0;
  padding: 0;
}

table {
  border-collapse: collapse;
  border-spacing: 0;
}

fieldset,
img {
  border: 0;
}

address,
caption,
cite,
code,
dfn,
em,
i,
optgroup,
th,
var {
  font-style: normal;
  font-weight: 400;
}

ol,
ul {
  list-style: none;
}

caption,
th {
  text-align: left;
}

h1,
h2,
h3,
h4,
h5,
h6 {
  font-size: 100%;
  font-weight: 400;
}

q:after,
q:before {
  content: '';
}

abbr,
acronym {
  border: 0;
  font-variant: normal;
}

sup {
  vertical-align: text-top;
}

sub {
  vertical-align: text-bottom;
}

button,
input,
optgroup,
option,
select,
textarea {
  font-family: inherit;
  font-size: inherit;
  font-style: inherit;
  font-weight: inherit;
}

button,
input,
select,
textarea {
  *font-size: 100%;
  line-height: 1.2;
}

input,
select,
textarea {
  outline: 0;
}

legend {
  color: #000;
}

body {
    font: 14px/1.4 "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Noto Sans CJK SC", "WenQuanYi Micro Hei", Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #151515;
    background-color: #fff;
}

.fr {
    float: right;
}

.full {
    width: 100%;
}

.fl {
    float: left;
}

.cl:after {
    content: "";
    display: block;
    visibility: hidden;
    clear: both;
}

textarea {
    background-color: white;
}

input {
    background-color: white;
}
</style>
