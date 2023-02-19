<script setup>
import {reactive} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()
const state = reactive({
  parcelNum: null
})

function go() {
  router.push(`/status/${state.parcelNum}`)
}
</script>

<template>
  <h1>Изменить статус посылки</h1>
  Введите номер посылки <br>
  <input type="number" v-model="state.parcelNum">
  <button @click="go">Далее</button>
</template>
