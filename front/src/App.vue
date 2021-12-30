<template>
  <v-app id="inspire">
    <v-app-bar
    app
    shrink-on-scroll
    color="white"
    elevation="15"
    dense
    >

      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title>Rand</v-toolbar-title>

      <v-spacer></v-spacer>
    

      <v-btn icon>
        <v-icon>mdi-dots-vertical</v-icon>

      </v-btn>
    </v-app-bar>

    <v-main>
      
    <router-view/>
      <v-container>
        <v-row>
          <div id="app">
        <div id="nav">
          
        </div>
      </div>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #d3f129;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
