<template>
  <div id="app">



    <message-box></message-box>
    <file-table></file-table>
    <Upload></Upload>

    <!--    <HelloWorld/>-->

  </div>
</template>

<script>
import messageBox from "@/components/messageBox";
import fileTable from "@/components/fileTable";
import Upload from "@/components/Upload";
export default {
  name: 'App',
  components: {
    Upload,
    // HelloWorld,
    fileTable,
    messageBox
  },
  data(){
    return {
    }

  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvtica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}


</style>
