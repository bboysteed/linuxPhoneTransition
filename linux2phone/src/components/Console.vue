<template>

  <div id="xterm" class="xterm" />

</template>


<script>
import 'xterm/css/xterm.css'
import { Terminal }  from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { AttachAddon } from 'xterm-addon-attach'

export default {
  name: 'Console',
  props: {
    socketURI: {
      type: String,
      default: ''
    },
  },
  mounted() {
    this.initSocket()
  },
  beforeDestroy() {
    this.socket.close()
    this.term.dispose()
  },
  methods: {
    initTerm() {
      const term = new Terminal({
        fontSize: 15,
        screenKeys: true,
        useStyle: true,
        cursorBlink: true,
      });
      const attachAddon = new AttachAddon(this.socket);
      const fitAddon = new FitAddon();

      term.loadAddon(attachAddon);
      term.loadAddon(fitAddon);
      term.open(document.getElementById('xterm'));
      fitAddon.fit();
      term.focus();
      this.term = term

      // term.onBinary(
      //     function (data) {
      //       let binData = new TextEncoder().encode(data)
      //       console.log(binData)
      //       this.socket.send(binData)
      //     }
      // )
      // term.onData(
      //     function (data) {
      //       this.socket.send(new TextEncoder().encode(data));
      //     }
      // )

    },
    initSocket() {
      this.socket = new WebSocket(this.socketURI);
      this.socketOnClose();
      this.socketOnOpen();
      this.socketOnError();
    },
    // socketOnMessage(){
    //   this.socket.onmessage = () =>{
    //
    //   }
    // },
    socketOnOpen() {
      this.socket.onopen = () => {
        // 链接成功后
        this.socket.binaryType = "arraybuffer";
        console.log('Console connect success!')

        this.initTerm()

      }
    },
    socketOnClose() {
      this.socket.onclose = () => {
        console.log('Console close socket')
      }
    },
    socketOnError() {
      this.socket.onerror = () => {
        console.log('Console socket 链接失败')
      }
    }
  }
}
</script>


