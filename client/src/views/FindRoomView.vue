<script setup lang="ts"></script>
<script lang="ts">
export default {
  data() {
    return {
      rooms: [
        { id: "1", diff: "easy" },
        { id: "2", diff: "high" },
      ],
    };
  },
};
</script>

<template>
  <main>
    <div v-for="room in rooms" v-bind:key="room.id">
      <div>
        <span>Room ID: {{ room.id }}</span>
      </div>
      <div>
        <span>Difficulty: {{ room.diff }} </span>
      </div>

      <button>Enter</button>
    </div>
  </main>
</template>

<style scoped lang="scss">
button {
  width: 50px;
  height: 50px;
}
</style>
