<template>
  <div class="overlay" :class="{ active: active }" @click="hideMenu"></div>
  <div class="menu" :class="{ active: active }">
    <div class="user">
      <div class="icon">
        <img :src="iconUrl" alt="avatar" class="icon-inner" />
      </div>
      <div class="name">
        <div class="username">{{ username }}</div>
        <div class="discrim">#{{ discrim }}</div>
      </div>
    </div>
    <div class="menu-grp large">
      <div class="menu-item" @click="move('/dashboard')">
        <fa-icon icon="bars" />Servers
      </div>
    </div>
    <div class="menu-grp">
      <!-- <div class="menu-item">Settings</div> -->
      <div class="menu-item" @click="logout">
        <fa-icon icon="right-from-bracket" />Log out
      </div>
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
      const user = this.userStore.user;
      return user ? user.username : "";
    },
    discrim(): string {
      const user = this.userStore.user;
      return user ? user.discriminator : "0000";
    },
    loggedIn(): boolean {
      return this.userStore.loggedIn;
    },
    active(): boolean {
      return this.navStore.active;
    },
    icon(): string {
      const user = this.userStore.user;
      return user ? user.avatar : "";
    },
    userId(): string {
      const user = this.userStore.user;
      return user ? user.id : "";
    },
    iconUrl(): string {
      if (!this.icon) {
        return `https://cdn.discordapp.com/embed/avatars/${
          parseInt(this.userId) % 5
        }.png`;
      }
      return this.icon.startsWith("a_")
        ? `https://cdn.discordapp.com/avatars/${this.userId}/${this.icon}.gif`
        : `https://cdn.discordapp.com/avatars/${this.userId}/${this.icon}.png`;
    },
  },
  methods: {
    hideMenu() {
      this.navStore.setActive(false);
    },
    move(to: string) {
      this.hideMenu();
      this.$router.push(to);
    },
    logout() {
      this.hideMenu();
      this.userStore.logOut();
      this.$router.push("/");
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
  gap: 1em;
  padding: 1.5em;
}

.menu-item {
  cursor: pointer;
  display: flex;
  font-size: 1rem;
  padding: 0.5em;
  border-radius: 0.5em;

  transition: 0.2s;
}

.menu-item:hover {
  background-color: rgba(0, 0, 0, 0.2);
}

.menu-item svg {
  font-size: 1.5rem;
  margin-right: 0.5em;
}

.user {
  display: flex;
  align-items: center;
  padding-bottom: 1em;
  border-bottom: 1px solid #aaa;
}

.icon {
  height: 48px;
  width: 48px;
}

.icon-inner {
  height: 100%;
  width: 100%;
  border-radius: 50%;
}

.name {
  margin-left: 0.5em;
  color: rgb(227, 227, 227);
}

.discrim {
  font-size: 0.7rem;
  color: rgb(174, 174, 174);
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
