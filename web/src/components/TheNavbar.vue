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

<script setup lang="ts">
import { useNavStore } from "@/stores/navbar";
import { computed } from "vue";
import { useRouter } from "vue-router";
import { useUserStore } from "../stores/user";

const userStore = useUserStore();
const navStore = useNavStore();
const router = useRouter();

const username = computed(() => {
  const user = userStore.user;
  return user ? user.username : "";
});

const discrim = computed(() => {
  const user = userStore.user;
  return user ? user.discriminator : "0000";
});

const loggedIn = computed(() => {
  return userStore.loggedIn;
});

const active = computed(() => {
  return navStore.active;
});

const icon = computed(() => {
  const user = userStore.user;
  return user ? user.avatar : "";
});

const userId = computed(() => {
  const user = userStore.user;
  return user ? user.id : "";
});

const iconUrl = computed(() => {
  if (!icon.value) {
    return `https://cdn.discordapp.com/embed/avatars/${
      parseInt(userId.value) % 5
    }.png`;
  }
  return icon.value.startsWith("a_")
    ? `https://cdn.discordapp.com/avatars/${userId.value}/${icon.value}.gif`
    : `https://cdn.discordapp.com/avatars/${userId.value}/${icon.value}.png`;
});

const hideMenu = () => {
  navStore.setActive(false);
};

const move = (to: string) => {
  hideMenu();
  router.push(to);
};

const logout = () => {
  hideMenu();
  userStore.logOut();
  router.push("/");
};
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
