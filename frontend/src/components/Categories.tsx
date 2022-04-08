import { memo, VFC } from 'react';

type Props = {
  title: string;
};

export const Categories: VFC<Props> = memo(({ title }) => {
  return <h2>Category: {title}</h2>;
});
