<script setup>
import {reactive} from "vue";
import $ from "jquery";

const state = reactive({
  username: null,
  password: null,
  isError: false,
  credentialsIsIncorrect: false
})

function authenticate() {
  $.ajax({
    url: "/api/auth/token",
    type: "POST",
    data: JSON.stringify({
      username: state.username,
      password: state.password
    }),
    success: (json) => {
      state.credentialsIsIncorrect = false
      window.localStorage.token = `Bearer ${json.token}`
      window.localStorage.username = json.user_nicename
      location.reload()
    },
    error: () => {
      state.credentialsIsIncorrect = true
    }
  })
}

function validateFields() {
  if (!state.username && !state.password) {
    state.isError = true
    return
  }
  state.isError = false
  authenticate()
}
</script>

<template>
  <table>
    <tr>
      <td>Имя пользователя</td>
      <td>
        <input type="text" v-model="state.username">
      </td>
    </tr>
    <tr>
      <td>Пароль</td>
      <td>
        <input type="password" v-model="state.password">
      </td>
    </tr>
  </table>
  <p v-if="state.isError" style="color: darkred">Все поля должны быть заполнены!</p>
  <p v-if="state.credentialsIsIncorrect" style="color: darkred">Имя пользователя или пароль неверны!</p>
  <button @click="validateFields">Войти</button>
</template>
