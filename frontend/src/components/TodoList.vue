<template>
  <ul style="display: flex;flex-direction: column">
    <li v-for="todo in todos" :key="todo.id" >
      {{ todo.task }}
    </li>
  </ul>
</template>

<script>
export default {
  props: ['todos'],
}
</script>