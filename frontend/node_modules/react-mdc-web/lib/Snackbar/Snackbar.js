'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var ROOT = 'mdc-snackbar';
var ACTIVE = ROOT + '--active';
var TEXT = ROOT + '__text';
var ACTION_WRAPPER = ROOT + '__action-wrapper';
var ACTION_BUTTON = ROOT + '__action-button';
var MULTILINE = ROOT + '--multiline';
var ACTION_ON_BOTTOM = ROOT + '--action-on-bottom';
var MESSAGE_TIMEOUT = 2750;

var Snackbar = function (_React$PureComponent) {
  _inherits(Snackbar, _React$PureComponent);

  function Snackbar(props) {
    _classCallCheck(this, Snackbar);

    var _this = _possibleConstructorReturn(this, (Snackbar.__proto__ || Object.getPrototypeOf(Snackbar)).call(this, props));

    _this.handleTimeout = _this.handleTimeout.bind(_this);
    return _this;
  }

  _createClass(Snackbar, [{
    key: 'componentDidMount',
    value: function componentDidMount() {
      var _props = this.props,
          open = _props.open,
          timeout = _props.timeout;

      if (open) {
        setTimeout(this.handleTimeout, timeout);
      }
    }
  }, {
    key: 'componentWillReceiveProps',
    value: function componentWillReceiveProps(_ref) {
      var nextOpen = _ref.open;
      var _props2 = this.props,
          open = _props2.open,
          timeout = _props2.timeout;

      if (nextOpen && !open) {
        setTimeout(this.handleTimeout, timeout);
      }
    }
  }, {
    key: 'handleTimeout',
    value: function handleTimeout() {
      var onTimeout = this.props.onTimeout;

      if (onTimeout) {
        onTimeout();
      }
    }
  }, {
    key: 'render',
    value: function render() {
      var _classnames;

      var _props3 = this.props,
          action = _props3.action,
          actionOnBottom = _props3.actionOnBottom,
          children = _props3.children,
          className = _props3.className,
          multiline = _props3.multiline,
          _onClick = _props3.onClick,
          onTimeout = _props3.onTimeout,
          open = _props3.open,
          timeout = _props3.timeout,
          otherProps = _objectWithoutProperties(_props3, ['action', 'actionOnBottom', 'children', 'className', 'multiline', 'onClick', 'onTimeout', 'open', 'timeout']);

      return _react2.default.createElement(
        'div',
        _extends({
          'aria-live': 'assertive',
          'aria-atomic': 'true',
          'aria-hidden': !open,
          className: (0, _classnames3.default)(ROOT, (_classnames = {}, _defineProperty(_classnames, ACTIVE, open), _defineProperty(_classnames, MULTILINE, multiline), _defineProperty(_classnames, ACTION_ON_BOTTOM, multiline && actionOnBottom), _classnames), className)
        }, otherProps),
        _react2.default.createElement(
          'div',
          { className: TEXT },
          children
        ),
        _react2.default.createElement(
          'div',
          {
            'aria-hidden': !_onClick || !action || !open,
            className: ACTION_WRAPPER
          },
          _react2.default.createElement(
            'button',
            {
              type: 'button',
              className: ACTION_BUTTON,
              onClick: function onClick(event) {
                if (_onClick) _onClick(event);
              }
            },
            action
          )
        )
      );
    }
  }]);

  return Snackbar;
}(_react2.default.PureComponent);

Snackbar.propTypes = {
  action: _propTypes2.default.string,
  actionOnBottom: _propTypes2.default.bool,
  children: _propTypes2.default.node,
  className: _propTypes2.default.string,
  multiline: _propTypes2.default.bool,
  onClick: _propTypes2.default.func,
  onTimeout: _propTypes2.default.func,
  open: _propTypes2.default.bool,
  timeout: _propTypes2.default.number
};
Snackbar.defaultProps = {
  timeout: MESSAGE_TIMEOUT
};
exports.default = Snackbar;