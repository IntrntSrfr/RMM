<template>
  <div class="guild">
    <router-link to="/dashboard">&lt;- Servers</router-link>
    <h1 v-if="currentGuild">{{ currentGuild.name }}</h1>
    <GuildMembers />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import GuildMembers from "@/components/GuildMembers.vue";
import { useGuildStore, type Guild } from "@/stores/guilds";

export default defineComponent({
  name: "TheGuild",
  components: { GuildMembers },
  setup() {
    const guildStore = useGuildStore();
    return { guildStore };
  },
  computed: {
    currentGuild(): Guild | undefined {
      const guildID = this.$route.params.guildID;
      if (!guildID) {
        return undefined;
      }
      return this.guildStore.getGuildByID(guildID.toString());
    },
  },
  async created() {
    const guildID = this.$route.params.guildID;
    if (!this.currentGuild) {
      await this.guildStore.fetchGuilds();
    }
    await this.guildStore.fetchGuildMembers(guildID.toString());
  },
});
</script>

<style scoped></style>
