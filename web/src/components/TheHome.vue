<template>
  <div>
    <h1>Recent Member Manager</h1>
    <p>Makes it easy to deal with large amount of recently joined users</p>
    <router-link v-if="loggedIn" to="/dashboard">Dashboard</router-link>
    <a
      v-else
      href="https://discord.com/api/oauth2/authorize?client_id=394162399348785152&redirect_uri=http%3A%2F%2Flocalhost%3A3000%2Fcallback&response_type=code&scope=identify%20guilds"
      >Sign in with Discord</a
    >
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useUserStore } from "@/stores/user";

export default defineComponent({
  name: "TheHome",
  setup() {
    const userStore = useUserStore();
    return { userStore };
  },
  computed: {
    loggedIn() {
      return this.userStore.loggedIn;
    },
  },
  created() {
    console.log(this.userStore.loggedIn);
  },
});
</script>

<style scoped></style>
