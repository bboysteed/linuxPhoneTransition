<template>
  <el-card
      :body-style="{ padding: '0px'}"
      shadow="hover"
      style="margin-top: 10px;border-radius: 10px"
  >
    <div v-for="tag in tags" :key="tag.ID" :class="tag.role==='pc'?'chat_item_left':'chat_item_right'">
      <img :src="tag.role==='pc'?pc_img:phone_img" alt="" class="role_img">
      <div :class="tag.role==='pc'?'content_box_left':'content_box_right'">
        {{ tag.content }}
      </div>
    </div>

    <el-input

        class="input_class"
        placeholder="请输入内容"
        v-model="textarea"
    >
      <el-button

          class="sub_button"
          slot="append"
          @click="send"
          type="primary"
          round
          size="small"
      >
        发送
      </el-button>
    </el-input>



  </el-card>
</template>

<script>
export default {
  name: "messageBox",
  data() {
    return {
      textarea: "",
      pc_img: require('@/assets/img.png'),
      phone_img: require('@/assets/img_1.png'),
      tags: [],
      wspath: "ws://" + window.location.href.split(":")[1] + ":9000/chatSocket",
      socket: "",
      role: "",
    }
  },
  mounted() {
    this.setRole()
    this.initws()
    this.getTags()
  },
  methods: {
    setRole() {
      let flag = navigator.userAgent.match(
          /(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i
      )
      if (flag) {
        this.role = "mobile"
      } else {
        this.role = "pc"
      }

    },
    initws() {
      if (typeof (WebSocket) === "undefined") {
        this.$message.error("您的浏览器不支持socket")

      } else {
        this.socket = new WebSocket(this.wspath)
        this.socket.onopen = this.open
        this.socket.onerror = this.error
        this.socket.onmessage = this.getMessage


      }
    },
    open: function () {
      console.log("socket连接成功")
    },
    error: function () {
      console.log("连接错误")
    },
    getMessage: function (msg) {
      let data = JSON.parse(msg.data)
      this.tags.push(data)
      console.log(this.tags)

    },
    send: function () {
      let data = {
        role: this.role,
        content: this.textarea,
      }
      if (data.content === "") {
        this.$message.error("请输入内容！")
        return
      }
      // this.tags.push(data)
      this.socket.send(JSON.stringify(data))
      this.textarea = ""
    },
    close: function () {
      console.log("socket已经关闭")
    },

    getTags() {
      this.$ajax.get(
          "/getMsgs",
          // this.$qs.stringify({
          //   role:"pc",
          //   content:this.textarea,
          // }),
      ).then(res => {
        console.log(res)
        this.tags = res.data
      }).catch(err => {
        console.log(err)
      })
    }
  }
}
</script>

<style>
.chat_item_left {
  max-width: 800px;
  /*width: 100%;*/
  padding: 10px 10px;
  display: flex;
  flex-wrap: wrap;
  overflow: visible;
}

.chat_item_right {
  max-width: 800px;
  /*width: 100%;*/
  padding: 10px 10px;
  flex-direction: row-reverse;
  display: flex;
  flex-wrap: wrap;
  overflow: visible;

}

.content_box_left {
  max-width: 290px;
  float: left;
  white-space: pre-wrap;
  text-align: left;
  position: relative;
  word-break: break-word;
  color: #8b4531;
  padding: 8px 10px;
  margin-left: 10px;
  background: #f5f6f7;
  border-radius: 5px
}

.content_box_right {
  max-width: 290px;
  white-space: pre-wrap;
  text-align: left;
  position: relative;
  word-break: break-word;
  color: #222226;
  padding: 8px 10px;
  margin-right: 10px;
  background: #cad9ff;
  border-radius: 5px
}


.role_img {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: block;
  float: left;
  cursor: pointer;
}

.input_class {

  padding: 10px 5px;
}
.sub_button{
  /*margin: 10px 10px;*/
  /*padding: 10px 10px;*/
}
</style>


