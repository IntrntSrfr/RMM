<template>
  <AppLoader v-if="guildsLoading" />
  <div v-else class="guild-list">
    <router-link
      class="list-item"
      v-for="(guild, i) in guilds"
      :to="{ name: 'guild', params: { guildID: guild.id } }"
      :key="i"
    >
      <GuildListItem :id="guild.id" :name="guild.name" :icon="guild.icon" />
    </router-link>
  </div>
  <MessageList />
</template>

<script setup lang="ts">
import { computed } from "vue";
import GuildListItem from "./GuildListItem.vue";
import { useGuildStore } from "@/stores/guilds";
import AppLoader from "./AppLoader.vue";
import MessageList from "./MessageList.vue";

const guildStore = useGuildStore();

const guilds = computed(() => {
  return guildStore.guilds;
});

const guildsLoading = computed(() => {
  return guildStore.loading;
});

const created = async () => {
  if (guilds.value.length) return;
  try {
    await guildStore.fetchGuilds();
  } catch (error) {
    console.log(error);
  }
};
await created();
</script>

<style scoped>
.guild-list {
  display: grid;
  grid-template-columns: 1fr;
}

.list-item {
  text-decoration: none;
  color: inherit;
}

.list-item + .list-item {
  border-top: 1px solid dodgerblue;
}

@media only screen and (min-width: 660px) {
  .guild-list {
    grid-template-columns: 1fr 1fr;
    gap: 1em;
  }

  .list-item + .list-item {
    border-top: none;
  }
}
</style>
