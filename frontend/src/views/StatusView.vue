<script setup>
import {useRoute} from "vue-router";
import {reactive} from "vue";
import $ from "jquery";

const route = useRoute()
const id = route.params.id
const state = reactive({
  message: 1,
  city: null,
  isCreated: false,
  isError: false
})

function save() {
  $.ajax({
    url: `/api/status/add/${id}`,
    type: "POST",
    headers: { Authorization: window.localStorage.token },
    data: JSON.stringify({
      message: state.message,
      city: state.city
    }),
    success: () => {
      state.isCreated = true
      state.isError = false
    },
    error: () => {
      state.isCreated = false
      state.isError = true
    }
  })
}
</script>

<template>
  <div>Номер посылки: {{id}}</div>

  <table>
    <tr>
      <td>Статус</td>
      <td>
        <select v-model.number="state.message">
          <option value="1">Присвоен трэк номер</option>
          <option value="2">Принято к перевозке</option>
          <option value="3">Покинуло место приема</option>
          <option value="4">Прибыло в пункт сортировки</option>
          <option value="5">Покинуло пункт сортировки</option>
          <option value="6">Прибыло в пункт выдачи</option>
          <option value="7">Выдано</option>
        </select>
      </td>
    </tr>
    <tr>
      <td>Город</td>
      <td>
        <input type="text" v-model="state.city">
      </td>
    </tr>
    <tr>
      <td>
        <button @click="save">Сохранить</button>
      </td>
    </tr>
  </table>

  <span v-if="state.isError">Произошла неизвестная ошибка. Возможно посыки не существует</span>
  <span v-if="state.isCreated">Статус обновлен</span>
</template>
