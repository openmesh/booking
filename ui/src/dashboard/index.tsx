import { Card, Col, Row, Space, Statistic, Typography } from 'antd';
import React from 'react';
import { Line, LineConfig } from '@ant-design/charts';
import { format } from 'date-fns';

export function Dashboard() {
  const collapsed = false;
  var config: LineConfig = {
    data: recentBookingsData,
    // height: 400,
    xField: 'date',
    yField: 'value',
    xAxis: {
      alias: 'Date',
    },
    yAxis: {
      alias: 'Value',
    },
    point: {
      shape: 'circle',
    },
    // point: {
    //   size: 5,
    //   shape: 'diamond',
    //   style: {
    //     fill: 'white',
    //     stroke: '#5B8FF9',
    //     lineWidth: 2,
    //   },
    // },
    tooltip: {
      showMarkers: false,
      title: data => data + 'wowow',
      formatter: data => {
        console.log(data);
        return { name: 'test', value: 'thing' };
      },
    },
    state: {
      active: {
        style: {
          shadowBlur: 4,
          stroke: '#000',
          fill: 'red',
        },
      },
    },
    interactions: [{ type: 'marker-active' }],
  };
  return (
    <>
      <Row wrap={true} gutter={24}>
        <Col className="flex-1">
          <Card>
            <Typography.Title level={4}>Recent Bookings</Typography.Title>
            <Space direction="vertical" className="w-full">
              <Statistic title="Booking value" value={112893} prefix="$" />
              <Statistic title="Booking quantity" value={45} />
              <Line {...config} className="mt-4" />
            </Space>
          </Card>
        </Col>
        <Col className="flex-1">
          <Card>
            <Typography.Title level={4}>Upcoming Bookings</Typography.Title>
          </Card>
        </Col>
      </Row>
    </>
  );
}

const recentBookingsData = [
  {
    date: format(new Date(2021, 9, 21), 'EEE dd'),
    value: 40,
    quantity: 3,
  },
  {
    date: format(new Date(2021, 9, 22), 'EEE dd'),
    value: 12,
    quantity: 1,
  },
  {
    date: format(new Date(2021, 9, 23), 'EEE dd'),
    value: 38,
    quantity: 6,
  },
  {
    date: format(new Date(2021, 9, 24), 'EEE dd'),
    value: 96,
    quantity: 9,
  },
  {
    date: format(new Date(2021, 9, 25), 'EEE dd'),
    value: 68,
    quantity: 4,
  },
  {
    date: format(new Date(2021, 9, 26), 'EEE dd'),
    value: 73,
    quantity: 2,
  },
  {
    date: format(new Date(2021, 9, 27), 'EEE dd'),
    value: 59,
    quantity: 7,
  },
];
