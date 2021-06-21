<template>
  <div>
    <el-card
        :body-style="{ padding: '0px'}"
        shadow="hover"
        style="margin-bottom: 40px"


    >

      <el-upload
          :action=uploadUrl
          drag
          :on-remove="handleRemove"
          :before-remove="beforeRemove"
          multiple
          :on-success="onUploadSuccess"
          :on-error="onUploadErr"
          :limit="10"
          :on-exceed="exceed"
          :file-list="fileList"
      >
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <div class="el-upload__tip" slot="tip"></div>

      </el-upload>
    </el-card>
  </div>

</template>

<script>


import axios from "axios";

export default {

  data() {
    return {
      uploadUrl:axios.defaults.baseURL+"upload",
      fileList: [],
      tableData:[]
    }
  },

  created() {

    },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList)

    },
    beforeRemove(file) {
      return this.$confirm(`确定移除 ${file.name}？`,'提示',{
        type: 'warning',
        center: true
      });

    },
    exceed(files, fileList) {
      this.$message.warning(`当前限制选择 3 个文件，本次选择了 ${files.length} 个文件，共选择了 ${files.length + fileList.length} 个文件`);


    },
    onUploadSuccess(response) {
      console.log(response)
      if (response.m === "ok") {

        this.$message.success("upload success!")

      }

    },
    onUploadErr() {
      this.$message.error("upload failed")

    },

  }

}
</script>

<style scoped>
.el-upload {
}
</style>