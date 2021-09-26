import {
  GithubOutlined,
  GoogleOutlined,
  TwitterOutlined,
} from '@ant-design/icons';
import { Button, Card, Col, Divider, Form, Input, Row } from 'antd';
import React from 'react';

export function Signin() {
  function handleGitHubButtonClick() {
    window.location.href = '/oauth/github';
  }

  return (
    <Row justify="center" align="middle">
      <Col flex="center" className="flex flex-col justify-center p-4">
        <Card className="mt-8">
          <Form layout="vertical">
            <Form.Item label="Email">
              <Input />
            </Form.Item>
            <Form.Item label="Password">
              <Input type="password" />
            </Form.Item>
            <Button className="w-full" type="primary">
              Sign in
            </Button>
          </Form>
          <Divider className="text-gray-600 text-xs my-4" plain>
            Or continue with
          </Divider>
          <Row gutter={8}>
            <Col span={8}>
              <Button icon={<GoogleOutlined />} className="w-full" />
            </Col>
            <Col span={8}>
              <Button
                icon={<GithubOutlined />}
                className="w-full"
                onClick={handleGitHubButtonClick}
              />
            </Col>
            <Col span={8}>
              <Button icon={<TwitterOutlined />} className="w-full" />
            </Col>
          </Row>
        </Card>
      </Col>
    </Row>
  );
}
