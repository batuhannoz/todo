<template>
  <li v-for="todo in todos" :key="todo.id">
    {{ todo.task }}
  </li>
</template>

<script>
export default {
  props: ['todos'],
}
</script>