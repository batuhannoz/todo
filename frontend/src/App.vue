<template>
  <AddTodo @addTask="addTask" />
  <TodoList :todos="todos"/>
</template>

<script>
import AddTodo from "./components/AddTodo.vue";
import TodoList from "./components/TodoList.vue";
import {addTodo} from "./api/todo";

export default {
  components: { AddTodo, TodoList },
  data() {
    return {
      todos: []
    };
  },
  methods: {
    addTask(task) {
      addTodo("someurl", task)
      this.todos.push({
        "task": task
      })
    }
  }
}
</script>

<style scoped>

</style>
