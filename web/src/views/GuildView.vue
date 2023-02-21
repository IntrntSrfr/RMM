<script setup lang="ts">
import AppLoader from "@/components/AppLoader.vue";
import TheGuild from "@/components/TheGuild.vue";
import AppScrollButton from "../components/AppScrollButton.vue";
</script>

<template>
  <Suspense>
    <TheGuild :is="TheGuild" />
    <template #fallback>
      <AppLoader />
    </template>
  </Suspense>
  <AppScrollButton />
</template>

<style scoped></style>
