<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { reactive } from "vue";

const state = reactive({
  isAuthenticated: false
})

// window.localStorage.token = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vbG9jYWxob3N0IiwiaWF0IjoxNjc0NTg0NDQyLCJuYmYiOjE2NzQ1ODQ0NDIsImV4cCI6MTY3NTE4OTI0MiwiZGF0YSI6eyJ1c2VyIjp7ImlkIjoiMSJ9fX0.NfsJ1gRMjGks0Wyg-aQRFh6FwGBgZtib-UOhlZ5broM"
// window.localStorage.username = "root"

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
