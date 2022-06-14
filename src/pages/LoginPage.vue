<script setup lang="ts">
import { nextTick } from 'vue';
import { cloneDeep, setToken } from '@/utils/tools';
import { useRouter } from 'vue-router';
import { login } from '@/api/Account';

const $router = useRouter();

const from = $ref({
  Email: '',
  Password: '',
  resMsg: {},
});

const Submit = async (e: any) => {
  const Elm = e.target;
  Elm.disabled = true;
  const res = await login({
    Email: from.Email,
    Password: from.Password,
  });
  Elm.disabled = false;
  from.resMsg = res;

  if (res.Code > 0) {
    await setToken(res.Data.Token);
    $router.push('/about');
  }

  nextTick(() => {
    console.info(e);
    console.info(cloneDeep(from));
  });
};
</script>

<template>
  <main>
    <h1>登录</h1>
    <input type="text" v-model="from.Email" name="Email" placeholder="请输入邮箱" autoComplete="new-password" />
    <br />
    <input
      type="password"
      v-model="from.Password"
      name="Password"
      placeholder="请输入密码"
      autoComplete="new-password"
    />
    <br />
    <br />
    Email: {{ from.Email }}
    <br />
    <br />
    Password: {{ from.Password }}
    <br />
    <br />

    <button @click="Submit">提交</button>
    <br />

    返回信息：
    <div v-if="from.resMsg" class="resMsg">{{ JSON.stringify(from.resMsg) }}</div>
  </main>
</template>

<style lang="less" scoped>
.resMsg {
  margin: 12px;
  border: 10px solid gray;
  padding: 15px;
  &.success {
    border-color: green;
  }
  &.fail {
    border-color: red;
  }
}
</style>
