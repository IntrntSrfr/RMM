<template>
  <div id="home">
    <h1>Recent Member Manager</h1>
    <h2>Quickly and easily deal with raids or bots</h2>
    <router-link v-if="loggedIn" class="btn" to="/dashboard">
      Dashboard
    </router-link>
    <a
      v-else
      class="btn"
      href="https://discord.com/api/oauth2/authorize?client_id=394162399348785152&redirect_uri=http%3A%2F%2Flocalhost%3A3000%2Fcallback&response_type=code&scope=identify%20guilds"
    >
      Sign in with Discord
    </a>
    <p>© {{ new Date().getFullYear() }} IntrntSrfr, All rights reserved</p>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useUserStore } from "@/stores/user";

const userStore = useUserStore();
const loggedIn = computed(() => {
  return userStore.loggedIn;
});
</script>

<style scoped>
#home {
  text-align: center;
  margin-top: 2em;
}

h1 {
  font-size: 2.5rem;
  font-weight: 700;
}

h2 {
  margin-top: 1em;
}

p {
  text-align: center;
  position: fixed;
  bottom: 1em;
  left: 0;
  width: 100%;
}

#home .btn {
  margin-top: 2em;
}

.btn {
  display: inline-block;
  padding: 0.75em 1em;
  font-size: 1rem;
  text-decoration: none;
  color: white;
  background-color: #5865f2;
  border-radius: 0.5em;

  transition: 0.2s;
}
.btn:hover {
  background-color: rgb(121, 131, 245);
}
</style>
