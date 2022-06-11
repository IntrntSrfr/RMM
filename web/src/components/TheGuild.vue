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
    //const guildID = this.$route.params.guildID;
    if (!this.currentGuild) {
      await this.guildStore.fetchGuilds();
    }
  },
});
</script>

<style scoped>
.header {
  display: flex;
  align-items: center;
  margin-bottom: 2em;
}

.header a {
  padding: 1em;
  margin-right: 1em;
}

.name {
  font-size: 1.5rem;
  text-overflow: ellipsis;
  overflow: hidden;
}
h3 {
  margin-bottom: 0.5em;
}
</style>
