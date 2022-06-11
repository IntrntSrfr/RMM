<template>
  <div class="overlay" :class="{ active: active }" @click="hideMenu"></div>
  <div class="menu" :class="{ active: active }">
    <div class="user">
      {{ username }}
    </div>
    <div class="menu-grp large">
      <div class="menu-item" @click="move('/dashboard')">Servers</div>
    </div>
    <div class="menu-grp">
      <div class="menu-item">Settings</div>
      <div class="menu-item">Log out</div>
    </div>
  </div>
</template>

<script lang="ts">
import { useNavStore } from "@/stores/navbar";
import { defineComponent } from "vue";
import { useUserStore } from "../stores/user";

export default defineComponent({
  name: "TheNavbar",
  setup() {
    const userStore = useUserStore();
    const navStore = useNavStore();
    return { userStore, navStore };
  },
  computed: {
    username(): string {
      if (!this.userStore.user) {
        return "";
      }
      return this.userStore.user.username;
    },
    loggedIn(): boolean {
      return this.userStore.loggedIn;
    },
    active(): boolean {
      return this.navStore.active;
    },
  },
  methods: {
    hideMenu() {
      this.navStore.setActive(false);
    },
    move(to: string) {
      this.$router.push(to);
      this.hideMenu();
    },
  },
});
</script>

<style scoped>
.menu {
  position: fixed;
  top: 0;
  left: -250px;
  height: 100%;
  width: 250px;

  background-color: #1b1e26;

  transition: 0.2s;
  visibility: hidden;
  z-index: 1;
}

.menu {
  display: flex;
  flex-direction: column;
  padding: 2em;
}

.menu-item {
  font-size: 1rem;
}

.large {
  flex-grow: 1;
}

.overlay {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  opacity: 0;
  visibility: hidden;
  overflow: hidden;
  background: black;
  transition: 0.2s;
  z-index: 1;
}

.active {
  visibility: visible;
}

.menu.active {
  left: 0;
}

.overlay.active {
  opacity: 0.2;
}
</style>
