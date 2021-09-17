<template>
  <div id="app">



    <message-box></message-box>
    <file-table></file-table>
    <Upload></Upload>

    <section>
      <div id="log" style="margin-top:20px;">
        <Console :socketURI="socketURI"></Console>
      </div>
    </section>


      <!--    <HelloWorld/>-->

  </div>
</template>

<script>
import messageBox from "@/components/messageBox";
import fileTable from "@/components/fileTable";
import Upload from "@/components/Upload";
import Console from "@/components/Console";
export default {
  name: 'App',
  components: {
    Console,
    Upload,
    // HelloWorld,
    fileTable,
    messageBox
  },
  data(){
    return {
      socketURI:'ws://127.0.0.1:10000/console'


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
