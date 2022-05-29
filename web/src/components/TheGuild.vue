<template>
  <h1>{{ currentGuild.name }}</h1>
  <GuildMembers />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useUserStore } from "@/stores/user";
import GuildMembers from "@/components/GuildMembers.vue";

export default defineComponent({
  name: "TheGuild",
  components: { GuildMembers },
  setup() {
    const userStore = useUserStore();
    return { userStore };
  },
  computed: {
    currentGuild() {
      const guildID = this.$route.params.guildID;
      if (!guildID) {
        return;
      }
      return this.userStore.getGuildByID(guildID);
    },
  },
});
</script>

<style scoped></style>
