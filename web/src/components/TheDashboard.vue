<template>
  <div class="dashboard">
    <div class="header">
      <fa-icon icon="bars" @click="openMenu" />
      <div>Servers</div>
    </div>
    <GuildList />
  </div>
</template>

<script setup lang="ts">
import { useNavStore } from "@/stores/navbar";
import GuildList from "./GuildList.vue";

const navStore = useNavStore();
const openMenu = () => {
  navStore.setActive(true);
};
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
}
.header {
  display: flex;
  align-items: center;
  margin-bottom: 1em;
}

.header svg {
  height: 24px;
  width: 24px;
  padding: 0.5em;
  cursor: pointer;
}

.header div {
  font-size: 1.5rem;
  margin-left: 0.25em;
}
</style>
