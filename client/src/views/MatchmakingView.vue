<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
  <main>
    <RouterLink to="/matchmaking/explore">find match</RouterLink>
    <button>create match</button>
  </main>
</template>

<style scoped lang="scss">
button {
  width: 50px;
  height: 50px;
}
</style>
