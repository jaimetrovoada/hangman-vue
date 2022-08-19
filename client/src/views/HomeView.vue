<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>
<script lang="ts">
let res: unknown;
export default {
  data() {
    return {
      response: res,
    };
  },
  methods: {
    pingSocket() {
      const url = "ws://localhost:8080/v1/api/ws";
      const conn = new WebSocket(url);
      const send = () => conn.send("hi");
      conn.onmessage = (msg) => {
        this.response = msg.data;
        console.log(msg);
      };
      try {
        conn.onopen = () => {
          send();
        };
      } catch (err) {
        console.log({ err });
      }
    },
  },
};
</script>

<template>
  <main>
    <TheWelcome />
    <RouterLink to="matchmaking">Find a match</RouterLink>
    <button @click="pingSocket">ping</button>
    <div>WebSocket message = {{ response }}</div>
  </main>
</template>
