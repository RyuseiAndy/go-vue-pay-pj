<template>
<div >
  <ul>
    <li v-for="item in items" :key="item.ID" @click="pageto(item.ID)">
      <item-card :item="item"></item-card>
    </li>
  </ul>
</div>
</template>


<script>
import axios from 'axios'
import ItemCard from './ItemCard'
export default {
  name: 'Home',
  components: {
    ItemCard
  },
  data () {
    return {
      items: []
    }
  },
  methods: {
    pageto: function (id) {
      this.$router.push(`/items/${id}`)
    }
  },
  created () {
    axios.get('http://localhost:8888/api/v1/items').then(res => {
      this.items = res.data
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.home{
  background: #24c6dc; /* fallback for old browsers */
  background: linear-gradient(to right, #514a9d 0%, #24c6dc 100%);
  opacity: 0.5;
  animation: huerotator 20s infinite alternate;
  color: cornsilk;
}

h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>