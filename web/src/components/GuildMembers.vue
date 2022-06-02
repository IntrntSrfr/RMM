<template>
  <div class="options">
    <AppButton text="Select all" @click="markAll" />
    <AppButton text="Select none" @click="markNone" />
    <AppButton text="Ban selected" @click="banSelected" />
  </div>
  <table>
    <thead>
      <tr>
        <th>
          <input
            type="checkbox"
            :checked="selectedMembers.length === members.length"
            :indeterminate="
              selectedMembers.length &&
              selectedMembers.length !== members.length
            "
            @click="selectedMembers.length ? markAll() : markNone()"
          />
        </th>
        <th>#</th>
        <th>ID</th>
        <th>Username</th>
        <th>Joined</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(m, i) in guildMembers" :key="i">
        <td>
          <input type="checkbox" v-model="m.checked" />
        </td>
        <td>{{ i }}</td>
        <td>{{ m.member.user.id }}</td>
        <td>
          {{ `${m.member.user.username}#${m.member.user.discriminator}` }}
        </td>
        <td>{{ timeAgo.format(m.joined) }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { Member } from "@/stores/user";
import AppButton from "@/components/AppButton.vue";
import TimeAgo from "javascript-time-ago";
import http from "@/http";
import { useGuildStore } from "@/stores/guilds";

type TableMember = {
  member: Member;
  checked: boolean;
  joined: Date;
};

export default defineComponent({
  name: "GuildMembers",
  components: { AppButton },
  setup() {
    const guildStore = useGuildStore();
    const timeAgo = new TimeAgo("en-US");
    return { guildStore, timeAgo };
  },
  data() {
    return {
      members: [] as TableMember[],
    };
  },
  computed: {
    selectedMembers(): TableMember[] {
      return this.members.filter((m: TableMember) => m.checked);
    },
    guildMembers(): TableMember[] {
      const guildID = this.$route.params.guildID;
      const g = this.guildStore.getGuildByID(guildID.toString());
      if (!g || !g.members) {
        return [];
      }

      return g.members
        .map((m: Member) => {
          return {
            member: m,
            checked: false,
            joined: new Date(m.joined_at),
          } as TableMember;
        })
        .sort((a: TableMember, b: TableMember) => {
          return b.joined.getTime() - a.joined.getTime();
        });
    },
  },
  async created() {
    const guildID = this.$route.params.guildID;
    if (!guildID) {
      return;
    }

    await this.guildStore.fetchGuildMembers(guildID.toString());
  },

  methods: {
    markAll() {
      this.members.forEach((m: TableMember) => (m.checked = true));
    },
    markNone(): void {
      this.members.forEach((m: TableMember) => (m.checked = false));
    },
    banSelected() {
      const ans = confirm("Are you SURE you want to ban ALL selected users?");
      if (!ans) {
        return;
      }

      console.log(this.selectedMembers);
      console.log(this.selectedMembers.map((m) => m.member.user.id));
    },
  },
});
</script>

<style scoped>
.options {
  display: flex;
  gap: 1em;
}

table {
  table-layout: fixed;
  width: 100%;
  border-collapse: collapse;
  border: 1px solid #a0a0a0;
}

thead {
  text-align: left;
  border-bottom: 1px solid var(--color-border);
}

thead th:nth-child(1) {
  width: 5%;
}

thead th:nth-child(2) {
  width: 5%;
}

thead th:nth-child(3) {
  width: 20%;
}

thead th:nth-child(4) {
  width: 40%;
}

thead th:nth-child(5) {
  width: 20%;
}

th,
td {
  padding: 0.5em;
}

</style>
