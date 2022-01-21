<template>
  <div>
    <v-card>
      <v-layout justify-center>
      <v-img
        height="280"
        width="210"
        src="https://cdn.vuetifyjs.com/images/cards/cooking.png"
      ></v-img>
      </v-layout>
      <v-card-title class="my-9">{{ item.Name }}</v-card-title>
      <v-card-text>{{ item.Description }}</v-card-text>
      <p class="align-right">{{ item.Amount }}円</p>
    </v-card>
    <payjp-checkout
      api-key="pk_test_892373dd331d28fa5152438e"
      client-id="d3d774f50bb006c26bac19402f0140a7228f8522"
      text="カードを情報を入力して購入"
      submit-text="購入確定"
      name-placeholder="田中 太郎"
      v-on:created="onTokenCreated"
      v-on:failed="onTokenFailed">
    </payjp-checkout>
    
    <p>{{ message }}</p>
    <v-btn depressed elevation="2" outlined>
    <router-link to="/">HOME</router-link>
    </v-btn>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'ItemCard',
  data () {
    return {
      item: {},
      message: ''
    }
  },
  created () {
    axios.get(`http://localhost:8888/api/v1/items/${this.$route.params.id}`).then(res => {
      this.item = res.data
    })
  },
  beforeDestroy () {
    window.PayjpCheckout = null
  },
  methods: {
    onTokenCreated: function (res) {
      console.log(res.id)
      const data = {Token: res.id}
      axios.post(`http://localhost:8888/api/v1/charge/items/${this.$route.params.id}`, data).then(res => {
        this.message = '商品の購入が完了しました'
      })
    },
    onTokenFailed: function (status, err) {
      console.log(status)
      console.log(err)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .align-right{
    text-align: middle;
  }
</style>