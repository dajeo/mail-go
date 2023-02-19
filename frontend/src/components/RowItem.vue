<script setup>
import {reactive} from "vue";
import $ from "jquery";

const props = defineProps(['id', 'num', 'name', 'unit', 'quantity', 'timestamp'])
const state = reactive({
  name: props.name,
  unit: props.unit,
  quantity: props.quantity,
  timestamp: new Date(props.timestamp)
})

function update() {
  $.ajax({
    url: `/api/items/update/${props.id}`,
    type: "POST",
    headers: { Authorization: window.localStorage.token },
    data: JSON.stringify({
      name: state.name,
      unit: state.unit,
      quantity: state.quantity
    })
  })
}
</script>

<template>
  <td>{{ props.num }}</td>
  <td>
    <input type="text" v-model="state.name" @change="update">
  </td>
  <td>
    <select v-model="state.unit" @change="update">
      <option>блок</option>
      <option>шт.</option>
      <option>тюбик</option>
      <option>пара</option>
      <option>балон</option>
    </select>
  </td>
  <td>
    <input type="number" min="1" v-model="state.quantity" @change="update" style="width: 45px">
  </td>
  <td>
    {{ state.timestamp.toLocaleDateString("ru-RU", {hour: "numeric", minute: "numeric"}) }}
  </td>
</template>
