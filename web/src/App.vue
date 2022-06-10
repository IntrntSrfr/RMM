<template>
  <RouterView />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useUserStore } from "./stores/user";

export default defineComponent({
  setup() {
    const userStore = useUserStore();

    return { userStore };
  },
  async created() {
    await this.userStore.fetchUser();
  },
});
</script>

<style>
@import "@/assets/base.css";

#app {
  max-width: 900px;
  margin: 0 auto;
  padding: 1.5rem;

  font-weight: normal;
}

@media only screen and (min-width: 900px) {
  #app {
    padding: 2rem;
  }
}
</style>
