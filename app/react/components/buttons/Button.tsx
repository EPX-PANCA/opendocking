import {
  AriaAttributes,
  ComponentType,
  forwardRef,
  MouseEventHandler,
  PropsWithChildren,
  ReactNode,
} from 'react';
import clsx from 'clsx';
import { motion } from 'framer-motion';

import { AutomationTestingProps } from '@/types';

import { Icon } from '@@/Icon';
import './Button.css';

type Type = 'submit' | 'button' | 'reset';
type Color =
  | 'default'
  | 'primary'
  | 'secondary'
  | 'danger'
  | 'link'
  | 'light'
  | 'dangerlight'
  | 'warninglight'
  | 'warning'
  | 'success'
  | 'blue'
  | 'none';
type Size = 'xsmall' | 'small' | 'medium' | 'large';

export interface Props<TasProps = unknown>
  extends AriaAttributes, AutomationTestingProps {
  icon?: ReactNode | ComponentType<unknown>;

  color?: Color;
  size?: Size;
  disabled?: boolean;
  title?: string;
  className?: string;
  type?: Type;
  as?: ComponentType<TasProps> | string;
  onClick?: MouseEventHandler<HTMLButtonElement>;
  mRef?: React.ForwardedRef<HTMLButtonElement>;
  props?: Omit<TasProps, keyof Props>;
}

export const ButtonWithRef = forwardRef<HTMLButtonElement, Omit<Props, 'mRef'>>(
  (props, ref) => (
    // eslint-disable-next-line react/jsx-props-no-spreading
    <Button {...props} mRef={ref} />
  )
);

ButtonWithRef.displayName = 'ButtonWithRef';

export function Button<TasProps = unknown>({
  type = 'button',
  color = 'primary',
  size = 'small',
  disabled = false,
  className,
  onClick,
  title,
  icon,
  children,
  as = 'button',
  props,
  mRef,
  ...ariaProps
}: PropsWithChildren<Props<TasProps>>) {
  const isButton = as === 'button';
  const Component = isButton
    ? (motion.button as unknown as 'button')
    : (as as 'button');
  const motionProps = isButton
    ? { whileHover: { scale: 1.02 }, whileTap: { scale: 0.97 } }
    : {};

  return (
    <Component
      ref={mRef}
      type={type}
      disabled={disabled}
      className={clsx(`btn btn-${color}`, sizeClass(size), className, {
        disabled,
      })}
      onClick={(e: React.MouseEvent<HTMLButtonElement>) => {
        if (!disabled) {
          onClick?.(e);
        }
      }}
      title={title}
      // eslint-disable-next-line react/jsx-props-no-spreading
      {...ariaProps}
      // eslint-disable-next-line react/jsx-props-no-spreading
      {...props}
      // eslint-disable-next-line react/jsx-props-no-spreading
      {...motionProps}
    >
      {icon && <Icon icon={icon} size={getIconSize(size)} />}
      {children}
    </Component>
  );
}

function getIconSize(size: Size) {
  switch (size) {
    case 'xsmall':
      return 'xs';
    case 'medium':
      return 'md';
    case 'large':
      return 'lg';
    case 'small':
    default:
      return 'sm';
  }
}

function sizeClass(size?: Size) {
  switch (size) {
    case 'large':
      return 'btn-lg';
    case 'medium':
      return 'btn-md';
    case 'xsmall':
      return 'btn-xs';
    default:
      return 'btn-sm';
  }
}
