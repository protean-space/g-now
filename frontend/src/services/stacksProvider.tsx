// @ts-ignore
import type { StacksProvider as StacksProviderType } from '@stacks/connect';

// 値として StacksProvider を使用
const StacksProvider = require('@stacks/connect').StacksProvider;

// 型を使う場合
const provider: StacksProviderType = new StacksProvider();


export const initializeStacksProvider = () => {
  // grobal stacks provider check
  if (typeof window.StacksProvider === 'undefined') {
    window.StacksProvider = provider;
  }
};