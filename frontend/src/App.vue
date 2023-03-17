<template>
  <AddTodo @addTask="addTask" />
  <TodoList :todos="todos"/>
</template>

<script>
import AddTodo from "./components/AddTodo.vue";
import TodoList from "./components/TodoList.vue";

export default {
  components: { AddTodo, TodoList },
  data() {
    return {
      todos: []
    };
  },
  methods: {
    addTask(task) {
      // TODO request backend
      this.todos.push({
        "task": task
      })
    }
  }
}
</script>

<style scoped>

</style>
