<template>
  <RouterView />
</template>

<script setup lang="ts">
import { onMounted } from "vue";
import { useUserStore } from "./stores/user";

const userStore = useUserStore();

onMounted(async () => {
  await userStore.fetchUser();
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
