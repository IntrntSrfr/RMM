<template>
  <div class="guild">
    <div class="header">
      <router-link to="/dashboard"><fa-icon icon="angle-left" /></router-link>
      <div v-if="currentGuild" class="name">{{ currentGuild.name }}</div>
    </div>
    <div v-if="currentGuild" class="content">
      <h3>New members</h3>
      <GuildMembers />
    </div>
  </div>
</template>

<script setup lang="ts">
import GuildMembers from "@/components/GuildMembers.vue";
import { useGuildStore } from "@/stores/guilds";
import { computed } from "vue";
import { useRoute } from "vue-router";

const guildStore = useGuildStore();
const route = useRoute();

const currentGuild = computed(() => {
  const guildID = route.params.guildID;
  if (!guildID) {
    return null;
  }
  return guildStore.getGuildByID(guildID.toString());
});

const created = async () => {
  if (!currentGuild.value) {
    await guildStore.fetchGuilds();
  }
};
await created();
</script>

<style scoped>
.header {
  display: flex;
  align-items: center;
  margin-bottom: 1em;
}

.header a {
  display: flex;
  color: #fff;
  padding: 0.5em;
  transition: 0.2s;
}

.header a:hover {
  color: #aaa;
}

.name {
  font-size: 1.5rem;
  margin-left: 0.25em;
  text-overflow: ellipsis;
  overflow: hidden;
}

h3 {
  margin-bottom: 0.5em;
}

a svg {
  height: 24px;
  width: 24px;
}
</style>
