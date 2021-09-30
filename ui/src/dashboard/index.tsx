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
    seriesField: 'type',
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
      title: data => format(new Date(data), 'EEE dd'),
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
    date: new Date(2021, 9, 21).toUTCString(),
    value: 40,
    type: 'value',
  },
  {
    date: new Date(2021, 9, 21).toUTCString(),
    value: 3,
    type: 'quantity',
  },
  {
    date: new Date(2021, 9, 22).toUTCString(),
    value: 12,
    type: 'value',
  },
  {
    date: new Date(2021, 9, 22).toUTCString(),
    value: 7,
    type: 'quantity',
  },
  {
    date: new Date(2021, 9, 23).toUTCString(),
    value: 38,
    type: 'value',
  },
  {
    date: new Date(2021, 9, 23).toUTCString(),
    value: 6,
    type: 'quantity',
  },
  {
    date: new Date(2021, 9, 24).toUTCString(),
    value: 96,
    type: 'value',
  },
  {
    date: new Date(2021, 9, 24).toUTCString(),
    value: 9,
    type: 'quantity',
  },
  {
    date: new Date(2021, 9, 25).toUTCString(),
    value: 68,
    quantity: 4,
  },
  {
    date: new Date(2021, 9, 25).toUTCString(),
    value: 68,
    quantity: 4,
  },
  {
    date: new Date(2021, 9, 26).toUTCString(),
    value: 73,
    quantity: 2,
  },
  {
    date: new Date(2021, 9, 27).toUTCString(),
    value: 59,
    quantity: 7,
  },
];
