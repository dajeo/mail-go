<script setup>
import {reactive} from "vue";
import {useRoute} from "vue-router";
import $ from "jquery";
import RowItem from "../components/RowItem.vue";

const route = useRoute()
const id = route.params.id
const state = reactive({
  items: [],
  empty: false,
  limit: false,
  error: false
})

$.ajax({
  url: `/api/items/get/all/${id}`,
  headers: { Authorization: window.localStorage.token },
  success: (json) => {
    state.error = false
    state.items = json
    state.empty = json.length === 0;
  },
  error: () => {
    state.error = true
    state.empty = false
  }
})

function addItem() {
  if (state.items.length === 15) {
    state.limit = true
    return
  }

  $.ajax({
    url: `/api/items/create/${id}`,
    type: "POST",
    headers: { Authorization: window.localStorage.token },
    success: (json) => {
      state.items.push(json)
    }
  })
}
</script>

<template>
  <div class="head">
    <div>Номер посылки: {{id}}</div>
    <button @click="addItem">Добавить</button>
  </div>

  <table v-if="state.items.length !== 0">
    <tr>
      <th>№</th>
      <th>Наименование</th>
      <th>ед.изм</th>
      <th>кол-во</th>
      <th>Дата</th>
    </tr>

    <tr v-for="(item, index) in state.items">
      <RowItem
          :id="item.id"
          :num="index + 1"
          :name="item.name"
          :unit="item.unit"
          :quantity="item.quantity"
          :timestamp="item.timestamp"
      />
    </tr>
  </table>

  <span v-if="state.limit">Вы добавили максимальное количество предметов</span>
  <span v-if="state.error">Посылка не найдена</span>
  <span v-if="state.empty">Список предметов пуст</span>
</template>

<style scoped>
.head {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
