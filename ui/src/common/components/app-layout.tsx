import {
  CalendarOutlined,
  PieChartOutlined,
  SettingOutlined,
  TagOutlined,
} from '@ant-design/icons';
import { Layout, Menu } from 'antd';
import { SelectInfo } from 'rc-menu/lib/interface';
import React from 'react';
import { useHistory, useLocation } from 'react-router';
import logo from '../../static/images/logo.svg';

export function AppLayout({ children }: { children: JSX.Element }) {
  const location = useLocation();
  const history = useHistory();

  function handleMenuItemSelected({ key }: SelectInfo) {
    history.push(key);
  }
  return (
    <Layout className="min-h-screen">
      <Layout.Sider collapsible theme="light">
        <img src={logo} alt="OpenMesh logo" className="h-10 mx-auto my-4" />
        <Menu
          mode="inline"
          theme="light"
          selectedKeys={[location.pathname]}
          onSelect={handleMenuItemSelected}
        >
          <Menu.Item key="/dashboard" icon={<PieChartOutlined />}>
            Dashboard
          </Menu.Item>
          <Menu.Item key="/bookings" icon={<CalendarOutlined />}>
            Bookings
          </Menu.Item>
          <Menu.Item key="/resources" icon={<TagOutlined />}>
            Resources
          </Menu.Item>
          <Menu.Item key="/settings" icon={<SettingOutlined />}>
            Settings
          </Menu.Item>
        </Menu>
      </Layout.Sider>
      <Layout className="bg-gray-100">
        <Layout.Header className="bg-gray-100" />
        <Layout.Content className="p-8">{children}</Layout.Content>
        <Layout.Footer className="text-center">
          Booking by OpenMesh ©2021 Created by Jack Caldwell
        </Layout.Footer>
      </Layout>
    </Layout>
  );
}
