<template>
  <div class="guild-list">
    <GuildListItem
      v-for="(guild, i) in guilds"
      :key="i"
      :name="guild.name"
      :icon="guild.icon"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useUserStore } from "../stores/user";
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
});
</script>

<style scoped></style>
