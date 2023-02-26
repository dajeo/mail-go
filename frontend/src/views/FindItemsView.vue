<script setup>
import {reactive} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()
const state = reactive({
  parcelNum: null
})

function go() {
  router.push(`/items/${state.parcelNum}`)
}
</script>

<template>
  <h1>Изменить содержимое посылки</h1>
  Введите номер посылки <br>
  <input type="number" v-model="state.parcelNum">
  <button @click="go">Изменить</button>
</template>
