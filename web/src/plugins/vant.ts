import { App as VM } from 'vue';
import {
  Button,
  Cell,
  CountDown,
  Popup,
  Form,
  Field,
  Cascader,
  Checkbox,
  CheckboxGroup,
  Radio,
  RadioGroup,
  Switch,
  Stepper,
  DatetimePicker,
  Picker,
  Uploader,
  Empty,
  List,
  PullRefresh,
  Tab,
  Tabs,
  Tag,
  Icon,
  Overlay,
  Loading
} from 'vant';

const plugins = [
  Button,
  Cell,
  CountDown,
  Popup,
  Form,
  Field,
  Cascader,
  Checkbox,
  CheckboxGroup,
  Radio,
  RadioGroup,
  Switch,
  Stepper,
  DatetimePicker,
  Picker,
  Uploader,
  Empty,
  List,
  PullRefresh,
  Tab,
  Tabs,
  Tag,
  Icon,
  Overlay,
  Loading
];

export const vantPlugins = {
  install: function(vm: VM) {
    plugins.forEach((item) => {
      //console.log('vant===',item.name, item)
      vm.component(item.name, item);
    });
  }
};
