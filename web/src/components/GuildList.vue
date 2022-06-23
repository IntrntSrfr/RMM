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
</template>

<script lang="ts">
import { defineComponent } from "vue";
import GuildListItem from "./GuildListItem.vue";
import { useGuildStore } from "@/stores/guilds";
import type { Guild } from "@/stores/guilds";
import AppLoader from "./AppLoader.vue";

export default defineComponent({
  name: "GuildList",
  components: { GuildListItem, AppLoader },
  setup() {
    const guildStore = useGuildStore();
    return { guildStore };
  },
  computed: {
    guilds(): Guild[] {
      return this.guildStore.guilds;
    },
    guildsLoading(): boolean {
      return this.guildStore.loading;
    },
  },
  async created() {
    try {
      await this.guildStore.fetchGuilds();
    } catch (error) {
      console.log(error);
    }
  },
});
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
