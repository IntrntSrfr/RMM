<template>
  <div id="nav">
    <div class="nav-grp">
      <router-link class="nav-item" to="/">RMM</router-link>
    </div>
    <div class="nav-grp">
      <router-link v-if="loggedIn" class="nav-item" to="/dashboard"
        >Dashboard</router-link
      >
      <a
        v-else
        class="nav-item"
        href="https://discord.com/api/oauth2/authorize?client_id=394162399348785152&redirect_uri=http%3A%2F%2Flocalhost%3A3000%2Fcallback&response_type=code&scope=identify%20guilds"
        >Sign in</a
      >
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useUserStore } from "../stores/user";

export default defineComponent({
  name: "TheNavbar",
  setup() {
    const userStore = useUserStore();
    return { userStore };
  },
  computed: {
    loggedIn() {
      return this.userStore.loggedIn;
    },
  },
});
</script>

<style scoped>
#nav {
  display: flex;
  justify-content: space-between;
}

.nav-grp {
  display: flex;
}

.nav-item {
  text-decoration: none;
  font-size: 1rem;
  color: var(--color-text);
  padding: 0.75em 1em;
}

.nav-item.router-link-active {
  color: dodgerblue;
}
</style>
