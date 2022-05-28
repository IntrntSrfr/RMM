<template>
  <div class="guild-list">
    <router-link
      class="list-item"
      v-for="(guild, i) in guilds"
      :to="{ name: 'guild', params: { guildID: guild.id } }"
      :key="i"
    >
      <GuildListItem :name="guild.name" :icon="guild.icon" />
    </router-link>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useUserStore } from "@/stores/user";
import GuildListItem from "./GuildListItem.vue";

export default defineComponent({
  name: "GuildList",
  components: { GuildListItem },
  setup() {
    const userStore = useUserStore();
    return { userStore };
  },
  computed: {
    guilds() {
      return this.userStore.guilds;
    },
  },
  async created() {
    try {
      await this.userStore.fetchGuilds();
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

.list-item{
  text-decoration: none;
  color: inherit;
}

.list-item + .list-item {
  border-top: 1px solid dodgerblue;
}
</style>
