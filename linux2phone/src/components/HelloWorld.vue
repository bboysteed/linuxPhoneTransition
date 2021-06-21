<template>
  <div>
    <v-form>
      <v-file-input
          v-model="fileValue"
          prepend-icon="mdi-file"
          label="file"
          counter
          show-size
          class="file"
          :disabled="loading.uploadIsLoading"
          :loading="loading.uploadIsLoading"
      >

      </v-file-input>
      <v-btn
          @click="uploadFile"
          outlined
          style="margin-top: 10px"
      >upload
      </v-btn>
    </v-form>
<!--    <fileTable></fileTable>-->

  </div>
</template>

<script>
// import fileTable from "@/components/fileTable";

export default {
  components: {
    // fileTable
  },
  data() {
    return {
      loading: {
        uploadIsLoading: false,
      },
      fileValue: [],
    }
  },
  mounted() {

  },
  methods: {
    uploadFile() {
      // let url = window.location.href+"upload/"
      const url = "http://192.168.50.168:9000/upload"
      let formData = new FormData()
      formData.append('file', this.fileValue)
      let config = {
        headers: {
          "Content-Type": "multipart/form-data"
        }
      }
      console.log(this.fileValue)
      this.loading.uploadIsLoading = true
      this.$ajax.post(
          url,
          formData,
          config,
      ).then((response) => {
        console.log("response:")
        console.log(response)
        if (response.data.m==="ok"){
          this.$message.success("upload success")
        }
        this.loading.uploadIsLoading = false
      }).catch(err => {
            console.log(err)
            this.loading.uploadIsLoading = false

          }
      )

    },
    // setValue(){
    //   this.fileValue
    // }
  }

}
</script>

<style>
.file {
  /*display: block;*/
  margin: auto;
  width: 500px;

}

</style>
