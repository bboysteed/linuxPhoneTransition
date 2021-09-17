<template>
  <div >
    <el-card
        shadow="hover"
        :body-style="{ padding: '0px' }"
        style="margin-top: 10px;border-radius: 10px"
        >

    <el-table
        :data="tableData"
        style="width: 100%;display: block"
        align="center"
    >
      <el-table-column
          style="margin:0 auto;"
          label="fileName"
          min-width="180">
        <template slot-scope="scope">
          <i class="el-icon-document"></i>
          <span style="margin-left: 10px">{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="size"
          min-width="70"

      style="text-align: center">
        <template slot-scope="scope">
          <el-popover trigger="hover" placement="top">
            <p>name: {{ scope.row.name }}</p>
            <p>size: {{ scope.row.size }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium" type="success" style="min-width: 100px">{{ scope.row.size }}</el-tag>
            </div>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column
          label="edit"
          min-width="180"
      >
        <template slot-scope="scope">
          <el-button
              size="mini"
              type="primary"
              @click="handleDownload(scope.$index, scope.row)">
            <i class="el-icon-download"></i>
          </el-button>
          <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)">
            <i class="el-icon-close"></i>
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    </el-card>
  </div>

</template>

<script>

export default {
  // props:{
  //   tD:{
  //     type:[String],
  //     // required:true
  //   }
  // },
  data() {
    return {
      tableData: []
    }
  },
  created() {
    this.initTableData()
      // this.tableData = this.tD
      // console.log(this.tableData)
    },
  methods: {
    handleDownload(index, row) {
      console.log(index, row)
      let downUrl = "http://"+window.location.href.split("//")[1]+"download/"+row.name
      this.downloadFile(downUrl)
      // let param = this.$qs.stringify({
      //   fileName:row.name
      // })
      // this.$ajax.get(
      //     downUrl,
      //     // param,
      //     {responseType:'blob'}
      // ).then((res) => {
      //   console.log(res)
      //   this.downloadFile(row.name,res.data)
      // }).catch((err)=>{
      //   console.log(err)
      // })
    },
    handleDelete(index, row) {
      console.log(index, row)
      console.log(row.name)
    },

    downloadFile(downurl) {
      let url = downurl
      let link = document.createElement('a')
      link.style.display = 'none'
      link.href = url;
      document.body.appendChild(link)
      link.target = "_blank"
      link.click()
      document.body.removeChild(link)
    },
    initTableData(){
      // let url = window.location.href+"getFiles"
      let url = "/getFiles"
      this.$ajax.get(
          url,
          // window.location.href + ""
      ).then(res=>{
        console.log(res.data)
        this.tableData = res.data
      }).catch(err => {
        console.log(err)
      }).finally(()=>{

      })
    },

  }
}
</script>

<style>


</style>