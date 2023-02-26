<script setup>
import {reactive} from "vue";
import {useRouter} from "vue-router";
import $ from "jquery";

const router = useRouter()

const state = reactive({
  isCompany: true,
  senderName: null,
  senderPhone: null,
  senderCity: null,
  battalion: null,
  company: null,
  platoon: null,
  department: null,
  recipientName: null,
  isError: false,
  isInternalError: false,
  isLoaded: false,
  loads: {}
})

$.ajax({
  url: "/api/parcels/load",
  headers: { Authorization: window.localStorage.token },
  success: (json) => {
    state.loads = json
    state.isLoaded = true
  }
})

function create() {
  $.ajax({
    url: "/api/parcels/create",
    type: "POST",
    headers: { Authorization: window.localStorage.token },
    data: JSON.stringify({
      senderName: state.senderName,
      senderPhone: state.senderPhone,
      senderCity: state.senderCity,
      recipientName: state.recipientName,
      battalion: state.battalion,
      company: state.company,
      platoon: state.platoon,
      department: state.department
    }),
    success: (json) => {
      state.isInternalError = false
      router.push(`/items/${json.id}`)
    },
    error: () => {
      state.isInternalError = true
    }
  })
}

function validateFields() {
  if (state.isCompany) {
    if (state.senderName &&
        state.senderPhone &&
        state.senderCity &&
        state.recipientName &&
        state.battalion &&
        state.company &&
        state.platoon &&
        state.department) { state.isError = false; create(); }
    else state.isError = true
  } else {
    if (state.senderName &&
        state.senderPhone &&
        state.senderCity &&
        state.recipientName &&
        state.battalion &&
        state.company) { state.isError = false; create(); }
    else state.isError = true
  }
}
</script>

<template>
  <h1>Создать посылку</h1>

  <div v-if="state.isLoaded">
    <input type="radio" v-model="state.isCompany" :value="true">
    <label>Рота</label>
    <br>
    <input type="radio" v-model="state.isCompany" :value="false">
    <label>Спец</label>

    <br>
    <table>
      <tr>
        <td>Отправитель:</td>
      </tr>
      <tr>
        <td>ФИО</td>
        <td>
          <input type="text" v-model="state.senderName">
        </td>
      </tr>
      <tr>
        <td>Номер телефона</td>
        <td>
          <input type="tel" v-model="state.senderPhone">
        </td>
      </tr>
      <tr>
        <td>Город</td>
        <td>
          <select v-model="state.senderCity">
            <option v-for="item in state.loads.cities">{{ item.name }}</option>
          </select>
        </td>
      </tr>
    </table>

    <br>
    <table>
      <tr>
        <td>Получатель:</td>
      </tr>
      <tr>
        <td>ФИО</td>
        <td>
          <input type="text" v-model="state.recipientName">
        </td>
      </tr>
      <tr>
        <td>Батальон</td>
        <td>
          <select v-model="state.battalion">
            <option v-for="item in state.loads.battalions">{{ item.name }}</option>
          </select>
        </td>
      </tr>
      <tr>
        <td>Рота</td>
        <td>
          <select v-if="state.isCompany" v-model="state.company">
            <option>1</option>
            <option>2</option>
            <option>3</option>
            <option>4</option>
          </select>
          <select v-else v-model="state.company">
            <option v-for="item in state.loads.specialUnits">{{ item.name }}</option>
          </select>
        </td>
      </tr>
      <tr v-if="state.isCompany">
        <td>Взвод</td>
        <td>
          <select v-model="state.platoon">
            <option>1</option>
            <option>2</option>
            <option>3</option>
            <option>4</option>
          </select>
        </td>
      </tr>
      <tr v-if="state.isCompany">
        <td>Отделение</td>
        <td>
          <select v-model="state.department">
            <option>1</option>
            <option>2</option>
            <option>3</option>
            <option>4</option>
          </select>
        </td>
      </tr>
    </table>

    <p v-if="state.isError" style="color: darkred">Все поля должны быть заполнены!</p>
    <p v-if="state.isInternalError" style="color: darkred">Произошла внутреняя ошибка!</p>
    <button @click="validateFields">Далее</button>
  </div>
  <p v-else>Загрузка...</p>
</template>
