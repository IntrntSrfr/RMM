<template>
  <div>Redirecting...</div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { useUserStore } from "../stores/user";

const userStore = useUserStore();
const router = useRouter();
const route = useRoute();

const created = async () => {
  const code = route.query.code;
  if (!code) {
    router.push("/");
    return;
  }

  try {
    await userStore.oauth(code.toString());
    await userStore.fetchUser();
  } catch (error) {
    router.push("/");
    console.log(error);
  }
  router.push("/dashboard");
};

await created();
</script>

<style scoped></style>
