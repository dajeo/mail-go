<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { reactive } from "vue";

const state = reactive({
  isAuthenticated: false
})

const token = window.localStorage.token
const username = window.localStorage.username

if (token && username) {
  state.isAuthenticated = true
}

function logout() {
  if (!confirm("Вы дейстивтельно хотите выйти?")) return
  window.localStorage.removeItem("token")
  window.localStorage.removeItem("username")
  location.reload()
}
</script>

<template>
  <header>
    <nav>
      <RouterLink to="/">Поиск</RouterLink>

      <RouterLink to="/createParcel" v-if="state.isAuthenticated">Создать посылку</RouterLink>
      <RouterLink to="/findItems" v-if="state.isAuthenticated">Изменить содержимое посылки</RouterLink>
      <RouterLink to="/findStatus" v-if="state.isAuthenticated">Изменить статус посылки</RouterLink>
      <a href="#" @click="logout" v-if="state.isAuthenticated">Выйти</a>

      <RouterLink to="/auth" v-if="!state.isAuthenticated">Авторизоваться</RouterLink>
    </nav>
  </header>

  <RouterView />
</template>

<style scoped>
nav {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  padding-bottom: 1rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  color: white;
}

nav a {
  padding: 0 1rem;
}

nav a:hover {
  color: white;
}
</style>
