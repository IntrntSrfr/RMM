<template>
  <div>{{ currentGuild }}</div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useUserStore } from "../stores/user";

export default defineComponent({
  name: "TheGuild",
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
