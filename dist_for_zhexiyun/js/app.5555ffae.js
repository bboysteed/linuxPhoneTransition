(function(e){function t(t){for(var o,i,r=t[0],l=t[1],c=t[2],d=0,p=[];d<r.length;d++)i=r[d],Object.prototype.hasOwnProperty.call(a,i)&&a[i]&&p.push(a[i][0]),a[i]=0;for(o in l)Object.prototype.hasOwnProperty.call(l,o)&&(e[o]=l[o]);u&&u(t);while(p.length)p.shift()();return s.push.apply(s,c||[]),n()}function n(){for(var e,t=0;t<s.length;t++){for(var n=s[t],o=!0,r=1;r<n.length;r++){var l=n[r];0!==a[l]&&(o=!1)}o&&(s.splice(t--,1),e=i(i.s=n[0]))}return e}var o={},a={app:0},s=[];function i(t){if(o[t])return o[t].exports;var n=o[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,i),n.l=!0,n.exports}i.m=e,i.c=o,i.d=function(e,t,n){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(i.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)i.d(n,o,function(t){return e[t]}.bind(null,o));return n},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/";var r=window["webpackJsonp"]=window["webpackJsonp"]||[],l=r.push.bind(r);r.push=t,r=r.slice();for(var c=0;c<r.length;c++)t(r[c]);var u=l;s.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"034f":function(e,t,n){"use strict";n("85ec")},"0caf":function(e,t,n){e.exports=n.p+"img/mountain.dbb72cf0.jpeg"},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var o=n("2b0e"),a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("message-box"),n("file-table"),n("Upload")],1)},s=[],i=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{style:{backgroundImage:"url("+e.bgimg+")",backgroundRepeat:"repeat",backgroundSize:"100% 100%"}},[n("el-card",{staticClass:"chat-card",staticStyle:{"border-radius":"10px","background-color":"transparent"},attrs:{"body-style":{padding:"0px"},shadow:"hover"}},[n("div",{staticClass:"chatWrapper",attrs:{id:"chatwrapper_"}},e._l(e.tags,(function(t){return n("div",{key:t.ID,class:"pc"===t.role?"chat_item_left":"chat_item_right"},[n("img",{staticClass:"role_img",attrs:{src:"pc"===t.role?e.pc_img:e.phone_img,alt:""}}),n("div",{class:"pc"===t.role?"content_box_left":"content_box_right"},[e._v(" "+e._s(t.content)+" ")])])})),0),n("el-input",{ref:"inputref",staticClass:"input_class",attrs:{id:"inputchat",placeholder:"请输入内容"},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.send(t)}},model:{value:e.textarea,callback:function(t){e.textarea=t},expression:"textarea"}},[n("el-button",{staticClass:"sub_button",attrs:{slot:"append",type:"primary",round:"",size:"small"},on:{click:e.send},slot:"append"},[e._v(" 发送 ")])],1)],1)],1)},r=[],l=(n("ac1f"),n("1276"),n("466d"),{name:"messageBox",data:function(){return{bgimg:n("0caf"),textarea:"",pc_img:n("7e77"),phone_img:n("f85b"),tags:[],wspath:"ws://"+window.location.href.split(":")[1]+"chatSocket",socket:"",role:""}},mounted:function(){this.setRole(),this.initws(),this.getTags()},updated:function(){var e=this,t=document.getElementById("chatwrapper_");t.scrollTop=t.scrollHeight,this.$nextTick((function(){e.$refs.inputref.focus()}))},methods:{setRole:function(){var e=navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i);this.role=e?"mobile":"pc"},initws:function(){"undefined"===typeof WebSocket?this.$message.error("您的浏览器不支持socket"):(this.socket=new WebSocket(this.wspath),this.socket.onopen=this.open,this.socket.onerror=this.error,this.socket.onmessage=this.getMessage)},open:function(){console.log("socket连接成功")},error:function(){console.log("连接错误")},getMessage:function(e){var t=JSON.parse(e.data);this.tags.push(t)},send:function(){var e={role:this.role,content:this.textarea};""!==e.content?(this.socket.send(JSON.stringify(e)),this.textarea="",this.$refs.inputref.focus()):this.$message.error("请输入内容！")},close:function(){console.log("socket已经关闭")},getTags:function(){var e=this;this.$ajax.get("/getMsgs").then((function(t){console.log(t),e.tags=t.data})).catch((function(e){console.log(e)}))}}}),c=l,u=(n("d4c3"),n("2877")),d=Object(u["a"])(c,i,r,!1,null,null,null),p=d.exports,f=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("el-card",{staticStyle:{"margin-top":"10px"},attrs:{shadow:"hover","body-style":{padding:"0px"}}},[n("el-table",{staticStyle:{width:"100%",display:"block"},attrs:{data:e.tableData,align:"center"}},[n("el-table-column",{staticStyle:{margin:"0 auto"},attrs:{label:"fileName","min-width":"180"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("i",{staticClass:"el-icon-document"}),n("span",{staticStyle:{"margin-left":"10px"}},[e._v(e._s(t.row.name))])]}}])}),n("el-table-column",{staticStyle:{"text-align":"center"},attrs:{label:"size","min-width":"70"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-popover",{attrs:{trigger:"hover",placement:"top"}},[n("p",[e._v("name: "+e._s(t.row.name))]),n("p",[e._v("size: "+e._s(t.row.size))]),n("div",{staticClass:"name-wrapper",attrs:{slot:"reference"},slot:"reference"},[n("el-tag",{staticStyle:{"min-width":"100px"},attrs:{size:"medium",type:"success"}},[e._v(e._s(t.row.size))])],1)])]}}])}),n("el-table-column",{attrs:{label:"edit","min-width":"180"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button",{attrs:{size:"mini",type:"primary"},on:{click:function(n){return e.handleDownload(t.$index,t.row)}}},[n("i",{staticClass:"el-icon-download"})]),n("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(n){return e.handleDelete(t.$index,t.row)}}},[n("i",{staticClass:"el-icon-close"})])]}}])})],1)],1)],1)},h=[],g=(n("b0c0"),n("d3b7"),{data:function(){return{tableData:[]}},created:function(){this.initTableData()},methods:{handleDownload:function(e,t){console.log(e,t);var n="http://"+window.location.href.split(":")[1]+"download/"+t.name;this.downloadFile(n)},handleDelete:function(e,t){console.log(e,t),console.log(t.name)},downloadFile:function(e){var t=e,n=document.createElement("a");n.style.display="none",n.href=t,document.body.appendChild(n),n.target="_blank",n.click(),document.body.removeChild(n)},initTableData:function(){var e=this,t="/getFiles";this.$ajax.get(t).then((function(t){console.log(t.data),e.tableData=t.data})).catch((function(e){console.log(e)})).finally((function(){}))}}}),m=g,b=Object(u["a"])(m,f,h,!1,null,null,null),v=b.exports,y=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("el-card",{attrs:{"body-style":{padding:"0px"},shadow:"hover"}},[n("el-upload",{attrs:{action:e.uploadUrl,drag:"","on-remove":e.handleRemove,"before-remove":e.beforeRemove,multiple:"","on-success":e.onUploadSuccess,"on-error":e.onUploadErr,limit:10,"on-exceed":e.exceed,"file-list":e.fileList}},[n("i",{staticClass:"el-icon-upload"}),n("div",{staticClass:"el-upload__text"},[e._v("将文件拖到此处，或"),n("em",[e._v("点击上传")])]),n("div",{staticClass:"el-upload__tip",attrs:{slot:"tip"},slot:"tip"})])],1)],1)},_=[],w=(n("99af"),n("bc3a")),x=n.n(w),k={data:function(){return{uploadUrl:x.a.defaults.baseURL+"upload",fileList:[],tableData:[]}},created:function(){},methods:{handleRemove:function(e,t){console.log(e,t)},beforeRemove:function(e){return this.$confirm("确定移除 ".concat(e.name,"？"),"提示",{type:"warning",center:!0})},exceed:function(e,t){this.$message.warning("当前限制选择 3 个文件，本次选择了 ".concat(e.length," 个文件，共选择了 ").concat(e.length+t.length," 个文件"))},onUploadSuccess:function(e){console.log(e),"ok"===e.m&&this.$message.success("upload success!")},onUploadErr:function(){this.$message.error("upload failed")}}},S=k,O=(n("9474"),Object(u["a"])(S,y,_,!1,null,"7a204146",null)),$=O.exports,j={name:"App",components:{Upload:$,fileTable:v,messageBox:p},data:function(){return{}}},C=j,E=(n("034f"),Object(u["a"])(C,a,s,!1,null,null,null)),D=E.exports,U=n("5c96"),z=n.n(U),P=(n("0fae"),n("ce5b")),M=n.n(P);n("bf40");o["default"].use(M.a);var R={},T=new M.a(R),B=n("4328"),N=n.n(B),I=n("8f94"),J=n.n(I);n("a7be");o["default"].use(J.a),o["default"].prototype.$ajax=x.a,o["default"].prototype.$qs=N.a,o["default"].config.productionTip=!1,o["default"].use(z.a,{size:"small",zIndex:2e3}),x.a.defaults.baseURL="/",console.log(Object({NODE_ENV:"production",BASE_URL:"/"})),new o["default"]({vuetify:T,render:function(e){return e(D)}}).$mount("#app")},"78f4":function(e,t,n){},"7e77":function(e,t,n){e.exports=n.p+"img/img.4849284d.png"},"85ec":function(e,t,n){},"894f":function(e,t,n){},9474:function(e,t,n){"use strict";n("78f4")},d4c3:function(e,t,n){"use strict";n("894f")},f85b:function(e,t,n){e.exports=n.p+"img/img_1.5832b61f.png"}});
//# sourceMappingURL=app.5555ffae.js.map